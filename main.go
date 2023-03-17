package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	data := `{
        "Athlete": "DEMTSCHENKO, Albert",
        "City": "Turin",
        "Country": "RUS",
        "Discipline": "Luge",
        "Event": "Singles",
        "Gender": "Men",
        "Medal": "Silver",
        "Season": "winter",
        "Sport": "Luge",
        "Year": 2006
    }`
	req, err := http.NewRequest("POST", "http://localhost:4080/api/games3/_doc", strings.NewReader(data))
	if err != nil {
		fmt.Println(err.Error())
	}
	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	fmt.Println(strconv.Itoa(resp.StatusCode))

}
