package ews

import "net/http"

type ExchangeCredentials interface {
	PrepareWebRequest(req *http.Request)
}
