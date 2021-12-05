package ewsCredentials

import (
	"net/http"
)

type webCredentials struct {
	userName string
	password string
}

func NewWebCredentials(userName, password string) ExchangeCredentials {
	return &webCredentials{
		userName: userName,
		password: password,
	}
}

func (c *webCredentials) PrepareWebRequest(req *http.Request) {
	req.SetBasicAuth(c.userName, c.password)
}
