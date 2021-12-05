//nolint
package ews

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

// client is soap client
type client struct {
	url     string
	opts    *options
	headers []interface{}
}

// newClient creates new SOAP client instance
func newClient(url string, opt ...option) *client {
	opts := defaultOptions
	for _, o := range opt {
		o(&opts)
	}
	return &client{
		url:  url,
		opts: &opts,
	}
}

type options struct {
	tlsCfg              *tls.Config
	credentials         ExchangeCredentials
	timeout             time.Duration
	conTimeout          time.Duration
	tlsHandshakeTimeout time.Duration
	client              httpClient
	httpHeaders         map[string]string
}

var defaultOptions = options{
	timeout:             30 * time.Second,
	conTimeout:          90 * time.Second,
	tlsHandshakeTimeout: 15 * time.Second,
}

// HTTPClient is a client which can make HTTP requests
// An example implementation is net/http.Client
type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// An option sets options such as credentials, tls, etc.
type option func(*options)

func withCredentials(credentials ExchangeCredentials) option {
	return func(o *options) {
		o.credentials = credentials
	}
}

// addHeader adds envelope header
// For correct behavior, every header must contain a `XMLName` field.  Refer to #121 for details
func (s *client) addHeader(header interface{}) {
	s.headers = append(s.headers, header)
}

// setHeaders sets envelope headers, overwriting any existing headers.
// For correct behavior, every header must contain a `XMLName` field.  Refer to #121 for details
func (s *client) setHeaders(headers ...interface{}) {
	s.headers = headers
}

const XmlNsSoapEnv string = "http://schemas.xmlsoap.org/soap/envelope/"

func (s *client) call(ctx context.Context, soapAction string, request, response interface{}) error {
	// SOAP envelope capable of namespace prefixes
	envelope := SOAPEnvelope{
		XmlNS: XmlNsSoapEnv,
	}

	if s.headers != nil && len(s.headers) > 0 {
		envelope.Header = &SOAPHeader{
			Headers: s.headers,
		}
	}

	envelope.Body.Content = request
	buffer := new(bytes.Buffer)
	encoder := xml.NewEncoder(buffer)

	if err := encoder.Encode(envelope); err != nil {
		return err
	}

	if err := encoder.Flush(); err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", s.url, buffer)
	if err != nil {
		return err
	}

	if s.opts.credentials != nil {
		s.opts.credentials.PrepareWebRequest(req)
	}

	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	req.Header.Add("SOAPAction", soapAction)

	if s.opts.httpHeaders != nil {
		for k, v := range s.opts.httpHeaders {
			req.Header.Set(k, v)
		}
	}
	req.Close = true

	client := s.opts.client
	if client == nil {
		tr := &http.Transport{
			TLSClientConfig: s.opts.tlsCfg,
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				d := net.Dialer{Timeout: s.opts.timeout}
				return d.DialContext(ctx, network, addr)
			},
			TLSHandshakeTimeout: s.opts.tlsHandshakeTimeout,
		}
		client = &http.Client{Timeout: s.opts.conTimeout, Transport: tr}
		s.opts.client = client
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		body, _ := ioutil.ReadAll(res.Body)
		return &HTTPError{
			StatusCode:   res.StatusCode,
			ResponseBody: body,
		}
	}

	// xml Decoder cannot handle namespace prefixes (yet),
	// so we have to use a namespace-less response envelope
	respEnvelope := new(SOAPEnvelopeResponse)
	respEnvelope.Body = SOAPBodyResponse{
		Content: response,
		Fault: &SOAPFault{
			Detail: nil,
		},
	}

	dec := xml.NewDecoder(res.Body)
	if err := dec.Decode(respEnvelope); err != nil {
		return err
	}

	return respEnvelope.Body.ErrorFromFault()
}

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"soap:Envelope"`
	XmlNS   string   `xml:"xmlns:soap,attr"`

	Header *SOAPHeader
	Body   SOAPBody
}

type SOAPHeader struct {
	XMLName xml.Name `xml:"soap:Header"`

	Headers []interface{}
}

type SOAPBody struct {
	XMLName xml.Name `xml:"soap:Body"`

	Content interface{} `xml:",omitempty"`

	// faultOccurred indicates whether the XML body included a fault;
	// we cannot simply store SOAPFault as a pointer to indicate this, since
	// fault is initialized to non-nil with user-provided detail type.
	faultOccurred bool
	Fault         *SOAPFault `xml:",omitempty"`
}

func (b *SOAPBody) ErrorFromFault() error {
	if b.faultOccurred {
		return b.Fault
	}
	b.Fault = nil
	return nil
}

type SOAPEnvelopeResponse struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  *SOAPHeaderResponse
	Body    SOAPBodyResponse
}

type SOAPHeaderResponse struct {
	XMLName xml.Name `xml:"Header"`

	Headers []interface{}
}

type SOAPBodyResponse struct {
	XMLName xml.Name `xml:"Body"`

	Content interface{} `xml:",omitempty"`

	// faultOccurred indicates whether the XML body included a fault;
	// we cannot simply store SOAPFault as a pointer to indicate this, since
	// fault is initialized to non-nil with user-provided detail type.
	faultOccurred bool
	Fault         *SOAPFault `xml:",omitempty"`
}

func (b *SOAPBodyResponse) UnmarshalXML(d *xml.Decoder, _ xml.StartElement) error {
	if b.Content == nil {
		return xml.UnmarshalError("Content must be a pointer to a struct")
	}

	var token xml.Token
	var err error
	consumed := false

Loop:
	for {
		if token, err = d.Token(); err != nil {
			return err
		}

		if token == nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if consumed {
				return xml.UnmarshalError("Found multiple elements inside SOAP body; not wrapped-document/literal WS-I compliant")
			} else if se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" && se.Name.Local == "Fault" {
				b.Content = nil

				b.faultOccurred = true
				if err := d.DecodeElement(b.Fault, &se); err != nil {
					return err
				}

				consumed = true
			} else {
				if err := d.DecodeElement(b.Content, &se); err != nil {
					return err
				}

				consumed = true
			}
		case xml.EndElement:
			break Loop
		}
	}

	return nil
}

func (b *SOAPBodyResponse) ErrorFromFault() error {
	if b.faultOccurred {
		return b.Fault
	}
	b.Fault = nil
	return nil
}

type SOAPFault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`

	Code   string     `xml:"faultcode,omitempty"`
	String string     `xml:"faultstring,omitempty"`
	Actor  string     `xml:"faultactor,omitempty"`
	Detail FaultError `xml:"detail,omitempty"`
}

func (f *SOAPFault) Error() string {
	if f.Detail != nil && f.Detail.HasData() {
		return f.Detail.ErrorString()
	}
	return f.String
}

// TODO работает ли? нужно ли?

type FaultError interface {
	// ErrorString should return a short version of the detail as a string,
	// which will be used in place of <faultstring> for the error message.
	// Set "HasData()" to always return false if <faultstring> error
	// message is preferred.
	ErrorString() string
	// HasData indicates whether the composite fault contains any data.
	HasData() bool
}

// HTTPError is returned whenever the HTTP request to the server fails
type HTTPError struct {
	// StatusCode is the status code returned in the HTTP response
	StatusCode int
	// ResponseBody contains the body returned in the HTTP response
	ResponseBody []byte
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTP Status %d: %s", e.StatusCode, string(e.ResponseBody))
}
