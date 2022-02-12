package user

import (
	"net/url"
	"strings"

	"github.com/hasusuf/sftpgo-api/types"
)

func (c *client) GetUser(username string) (*types.User, error) {
	var localVarReturnValue *types.User

	localVarPath := "/api/v2/users/{username}"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", url.PathEscape(parameterToString(username, "")), -1)

	req := Request{
		Method:   GET,
		Endpoint: localVarPath,
	}

	// Make the API call
	localVarHTTPResponse, err := c.Call(req)
	if err != nil {
		return nil, err
	}

	// Bind the response body
	err = c.BindResponse(&localVarReturnValue, localVarHTTPResponse)
	if err != nil {
		return nil, err
	}

	return localVarReturnValue, nil
}
