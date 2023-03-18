package ioindexer

import (
	"fmt"
	"io/ioutil"
)

func check(er error) {
	if er != nil {
		panic(er)
	}
}

func GetDataBody(nameFolderData string) string {
	c, err := ioutil.ReadDir("./data/" + nameFolderData)
	check(err)

	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	return `{
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
}
