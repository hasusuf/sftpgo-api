package user

import (
	"strconv"

	"github.com/hasusuf/sftpgo-api/common"
	"github.com/hasusuf/sftpgo-api/types"
)

func (c *client) GetUsers(request types.GetUsersRequest) ([]types.User, error) {
	var localVarReturnValue []types.User

	queryParams := make(map[string]string)
	if request.Offset > 0 {
		queryParams["offset"] = strconv.Itoa(request.Offset)
	}
	if request.Limit > 0 {
		queryParams["limit"] = strconv.Itoa(request.Limit)
	}
	if request.Order != `` {
		queryParams["order"] = request.Order
	}

	req := common.Request{
		Method:      common.GET,
		Endpoint:    `/api/v2/users`,
		QueryParams: queryParams,
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
