package user

import (
	"github.com/hasusuf/sftpgo-api/common"
	"github.com/hasusuf/sftpgo-api/types"
)

func (c *client) CreateUser(request *types.User) (*types.User, error) {
	var localVarReturnValue *types.User

	localVarPath := "/api/v2/users"

	localVarRequestPayload, err := request.MarshalJSON()
	if err != nil {
		return nil, err
	}

	req := common.Request{
		Method:   common.POST,
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
