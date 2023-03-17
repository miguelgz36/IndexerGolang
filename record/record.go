package record

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/miguelgz36/IndexerGolang/connection"
)

func PostData(data string) {
	req, err := http.NewRequest("POST", "http://localhost:4080/api/games3/_doc", strings.NewReader(data))
	if err != nil {
		fmt.Println(err.Error())
	}

	connection.SetHeaders(req)

	doPost(req)
}

func doPost(req *http.Request) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	fmt.Println(strconv.Itoa(resp.StatusCode))
}
