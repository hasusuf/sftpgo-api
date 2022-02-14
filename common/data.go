package common

import (
	"net/http"
	"regexp"
)

// Method contains the supported HTTP verbs
type Method string

// Scheme contains the supported HTTP protocol schemes
type Scheme string

// Request holds the request to an API Call
type Request struct {
	Method      Method
	Url         string
	Host        string  // e.g. api.example.com
	Scheme      *Scheme // e.g. http || https
	Endpoint    string  // e.g. api/v2/token
	Headers     map[string]string
	QueryParams map[string]string
	Body        []byte
}

// Options allows modification of client headers and other settings
// See https://golang.org/pkg/net/http
type Options struct {
	HTTPClient  *http.Client
	AccessToken string
	Host        string
	Scheme      *Scheme
	Credentials Credentials
}

type Credentials struct {
	Username string
	Password string
}

// Supported HTTP verbs.
const (
	POST    Method = "POST"
	GET     Method = "GET"
	HEAD    Method = "HEAD"
	PUT     Method = "PUT"
	DELETE  Method = "DELETE"
	PATCH   Method = "PATCH"
	OPTIONS Method = "OPTIONS"
	HTTP    Scheme = "http"
	HTTPS   Scheme = "https"
)

var (
	JsonCheck = regexp.MustCompile(`(?i:(?:application|text)/(?:vnd\.[^;]+\+)?json)`)
	XmlCheck  = regexp.MustCompile(`(?i:(?:application|text)/xml)`)
)

// defaultClient is used if no custom HTTP client is defined
var defaultClient = &Options{HTTPClient: &http.Client{}}
