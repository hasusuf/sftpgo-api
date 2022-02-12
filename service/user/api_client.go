package user

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/hasusuf/sftpgo-api/types"
)

var (
	jsonCheck = regexp.MustCompile(`(?i:(?:application|text)/(?:vnd\.[^;]+\+)?json)`)
	xmlCheck  = regexp.MustCompile(`(?i:(?:application|text)/xml)`)
)

// Method contains the supported HTTP verbs
type Method string

// Scheme contains the supported HTTP protocol schemes
type Scheme string

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

// Request holds the request to an API Call
type Request struct {
	Method      Method
	url         string
	Host        string  // e.g. api.example.com
	Scheme      *Scheme // e.g. http || https
	Endpoint    string  // e.g. api/v2/token
	Headers     map[string]string
	QueryParams map[string]string
	Body        []byte
}

type Client interface {
	GetUsers(request types.GetUsersRequest) ([]types.User, error)
	GetUser(username string) (*types.User, error)
	CreateUser(request *types.User) (*types.User, error)
	UpdateUser(username string, request *types.User) (*types.User, error)
}

// DefaultClient is used if no custom HTTP client is defined
var DefaultClient = &client{HTTPClient: &http.Client{}}

// Client allows modification of client headers and other settings
// See https://golang.org/pkg/net/http
type client struct {
	HTTPClient  *http.Client
	AccessToken string
	Host        string
	Scheme      *Scheme
}

func New(host, username, password string, scheme Scheme) (*client, error) {
	var localVarReturnValue types.Token

	c := client{
		HTTPClient: &http.Client{Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: 5 * time.Second,
			}).DialContext,
		}},
	}

	// Build basic credentials
	data := []byte(fmt.Sprintf("%s:%s", username, password))
	token := base64.StdEncoding.EncodeToString(data)

	// Build the header
	Headers := make(map[string]string)
	Headers["Authorization"] = "Basic " + token

	// Build the request params
	request := Request{
		Method:   GET,
		Host:     host,
		Scheme:   &scheme,
		Endpoint: "/api/v2/token",
		Headers:  Headers,
	}

	// Make the API call
	localVarHTTPResponse, err := Call(request)
	if err != nil {
		return nil, err
	}

	err = c.BindResponse(&localVarReturnValue, localVarHTTPResponse)
	if err != nil {
		return nil, err
	}

	c.Host = request.Host
	c.AccessToken = "Bearer " + localVarReturnValue.GetAccessToken()
	c.Scheme = &scheme

	return &c, nil
}

// Call uses the DefaultClient to send your request
func Call(request Request) (*http.Response, error) {
	return SendWithContext(context.Background(), request)
}

// SendWithContext uses the DefaultClient to send your request with the provided context
func SendWithContext(ctx context.Context, request Request) (*http.Response, error) {
	return DefaultClient.SendWithContext(ctx, request)
}

// Call will build your request, make the request, and build your response
func (c *client) Call(request Request) (*http.Response, error) {
	return c.SendWithContext(context.Background(), request)
}

// SendWithContext will build your request passing in the provided context, make the request, and build your response
func (c *client) SendWithContext(ctx context.Context, request Request) (*http.Response, error) {
	// Build the HTTP request URL
	request.url = c.BuildURL(request)

	// Include the access token to HTTP request header
	if len(c.AccessToken) > 0 {
		if request.Headers == nil {
			request.Headers = make(map[string]string)
		}

		request.Headers["Authorization"] = c.AccessToken
	}

	// Build the HTTP request object
	req, err := BuildRequestObject(request)
	if err != nil {
		return nil, err
	}
	// Pass in the user provided context
	req = req.WithContext(ctx)

	// Build the HTTP client and make the request
	res, err := c.MakeRequest(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// BuildURL builds the request URL
func (c *client) BuildURL(request Request) string {
	uri := url.URL{}
	if request.Scheme != nil {
		uri.Scheme = string(*request.Scheme)
	} else {
		uri.Scheme = string(*c.Scheme)
	}

	if request.Host != `` {
		uri.Host = request.Host
	} else {
		uri.Host = c.Host
	}

	uri.Path = request.Endpoint

	return uri.String()
}

// MakeRequest makes the API call
func (c *client) MakeRequest(req *http.Request) (*http.Response, error) {
	localVarHTTPResponse, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return nil, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		var v types.ApiResponse
		err = c.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}

		if localVarHTTPResponse.StatusCode >= 401 {
			return nil, errors.New(localVarHTTPResponse.Status)
		}

		return nil, errors.New(v.GetError())
	}

	return localVarHTTPResponse, nil
}

// BuildRequestObject creates the HTTP request object
func BuildRequestObject(request Request) (*http.Request, error) {
	// Add any query parameters to the URL
	if len(request.QueryParams) != 0 {
		request.url = AddQueryParameters(request.url, request.QueryParams)
	}
	req, err := http.NewRequest(string(request.Method), request.url, bytes.NewBuffer(request.Body))
	if err != nil {
		return req, err
	}
	for key, value := range request.Headers {
		req.Header.Set(key, value)
	}
	_, exists := req.Header["Content-Type"]
	if len(request.Body) > 0 && !exists {
		req.Header.Set("Content-Type", "application/json")
	}
	return req, err
}

// AddQueryParameters adds query parameters to the URL
func AddQueryParameters(baseURL string, queryParams map[string]string) string {
	baseURL += "?"
	params := url.Values{}
	for key, value := range queryParams {
		params.Add(key, value)
	}
	return baseURL + params.Encode()
}

func (c *client) Decode(v interface{}, b []byte, contentType string) (err error) {
	if len(b) == 0 {
		return nil
	}
	if s, ok := v.(*string); ok {
		*s = string(b)
		return nil
	}
	if f, ok := v.(**os.File); ok {
		*f, err = ioutil.TempFile("", "HttpClientFile")
		if err != nil {
			return
		}
		_, err = (*f).Write(b)
		if err != nil {
			return
		}
		_, err = (*f).Seek(0, io.SeekStart)
		return
	}
	if xmlCheck.MatchString(contentType) {
		if err = xml.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	}
	if jsonCheck.MatchString(contentType) {
		if actualObj, ok := v.(interface{ GetActualInstance() interface{} }); ok { // oneOf, anyOf schemas
			if unmarshalObj, ok := actualObj.(interface{ UnmarshalJSON([]byte) error }); ok { // make sure it has UnmarshalJSON defined
				if err = unmarshalObj.UnmarshalJSON(b); err != nil {
					return err
				}
			} else {
				return errors.New("unknown type with GetActualInstance but no unmarshalObj.UnmarshalJSON defined")
			}
		} else if err = json.Unmarshal(b, v); err != nil { // simple model
			return err
		}
		return nil
	}
	return errors.New("undefined response type")
}

// parameterToString convert interface{} parameters to string, using a delimiter if format is provided.
func parameterToString(obj interface{}, collectionFormat string) string {
	var delimiter string

	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	case "csv":
		delimiter = ","
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
	} else if t, ok := obj.(time.Time); ok {
		return t.Format(time.RFC3339)
	}

	return fmt.Sprintf("%v", obj)
}

func (c *client) BindResponse(v interface{}, resp *http.Response) error {
	localVarBody, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))

	err = c.Decode(&v, localVarBody, resp.Header.Get("Content-Type"))
	if err != nil {
		return err
	}

	return nil
}
