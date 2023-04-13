package record

import (
	"net/http"
	"strings"

	"github.com/miguelgz36/IndexerGolang/connection"
	"github.com/miguelgz36/IndexerGolang/errors"
)

func PostData(data string) {

	req, err := http.NewRequest("POST", "http://localhost:4080/api/emails/_doc", strings.NewReader(data))
	errors.Check(err)

	connection.SetHeaders(req)

	doPost(req)
}

func doPost(req *http.Request) {
	resp, err := http.DefaultClient.Do(req)
	errors.Check(err)
	defer resp.Body.Close()
}
