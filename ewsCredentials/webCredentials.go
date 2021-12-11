package ewsCredentials

import (
	"net/http"

	"github.com/anfilat/go-ews/internal/base"
)

type webCredentials struct {
	userName string
	password string
}

func NewWebCredentials(userName, password string) base.ExchangeCredentials {
	return &webCredentials{
		userName: userName,
		password: password,
	}
}

func (c *webCredentials) PrepareWebRequest(req *http.Request) {
	req.SetBasicAuth(c.userName, c.password)
}
