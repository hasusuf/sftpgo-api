package common

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/hasusuf/sftpgo-api/types"
)

func ResolveAccessToken(options *Options) (*types.Token, error) {
	var localVarToken types.Token

	// Build basic credentials
	data := []byte(fmt.Sprintf("%s:%s", options.Credentials.Username, options.Credentials.Password))
	basicAuth := base64.StdEncoding.EncodeToString(data)

	// Build the header
	Headers := make(map[string]string)
	Headers["Authorization"] = "Basic " + basicAuth

	// Build the request params
	request := Request{
		Method:   GET,
		Host:     options.Host,
		Scheme:   options.Scheme,
		Endpoint: "/api/v2/token",
		Headers:  Headers,
	}

	// Make the API call
	localVarHTTPResponse, err := Call(request)
	if err != nil {
		return nil, err
	}

	err = options.BindResponse(&localVarToken, localVarHTTPResponse)
	if err != nil {
		return nil, err
	}

	options.Host = request.Host
	options.AccessToken = "Bearer " + localVarToken.GetAccessToken()

	return &localVarToken, nil
}

func ResolveHTTPClient(options *Options) {
	if options.HTTPClient == nil {
		options.HTTPClient = http.DefaultClient
	}
}
