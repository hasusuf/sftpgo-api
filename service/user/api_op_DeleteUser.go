package user

import (
	"net/url"
	"strings"

	"github.com/hasusuf/sftpgo-api/common"
	"github.com/hasusuf/sftpgo-api/types"
)

func (c *client) DeleteUser(username string, request *types.User) (*types.User, error) {
	var localVarReturnValue *types.User

	localVarPath := "/api/v2/users/{username}"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", url.PathEscape(common.ParameterToString(username, "")), -1)

	localVarRequestPayload, err := request.MarshalJSON()
	if err != nil {
		return nil, err
	}

	req := common.Request{
		Method:   common.PUT,
		Endpoint: localVarPath,
		Body:     localVarRequestPayload,
	}

	// Make the API call
	localVarHTTPResponse, err := c.options.Call(req)
	if err != nil {
		return nil, err
	}

	// Bind the response body
	err = c.options.BindResponse(&localVarReturnValue, localVarHTTPResponse)
	if err != nil {
		return nil, err
	}

	return localVarReturnValue, nil
}
