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

	resp, err := doPost(req)

	if resp != nil {
		resp.Body.Close()
	}
	if err != nil {
		PostData(data)
	}
}

func doPost(req *http.Request) (*http.Response, error) {
	resp, err := http.DefaultClient.Do(req)
	return resp, err
}
