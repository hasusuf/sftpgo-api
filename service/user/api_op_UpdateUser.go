package user

import (
	"net/url"
	"strings"

	"github.com/hasusuf/sftpgo-api/types"
)

func (c *client) UpdateUser(username string, request *types.User) (*types.User, error) {
	var localVarReturnValue *types.User

	localVarPath := "/api/v2/users/{username}"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", url.PathEscape(parameterToString(username, "")), -1)

	localVarRequestPayload, err := request.MarshalJSON()
	if err != nil {
		return nil, err
	}

	req := Request{
		Method:   PUT,
		Endpoint: localVarPath,
		Body:     localVarRequestPayload,
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
