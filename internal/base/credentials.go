package base

import "net/http"

type ExchangeCredentials interface {
	PrepareWebRequest(req *http.Request)
}
