package connection

import (
	"net/http"
)

type Credentials struct {
	user     string
	password string
}

const CONTENT_TYPE_NAME string = "Content-Type"
const USER_AGENT_NAME string = "User-Agent"

const CONTENT_TYPE_VALUE string = "application/json"
const USER_AGENT_VALUE string = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36"

func SetHeaders(req *http.Request) {
	credentials := getCredentials()
	req.SetBasicAuth(credentials.user, credentials.password)
	req.Header.Set(CONTENT_TYPE_NAME, CONTENT_TYPE_VALUE)
	req.Header.Set(USER_AGENT_NAME, USER_AGENT_VALUE)
}

func getCredentials() Credentials {
	return Credentials{
		user:     "admin",
		password: "Complexpass#123",
	}
}
