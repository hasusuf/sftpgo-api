package health

import "github.com/hasusuf/sftpgo-api/common"

func (c *client) Health() (ok bool) {
	var localVarReturnValue bool

	localVarPath := "/api/v2/health"

	req := common.Request{
		Method:   common.GET,
		Endpoint: localVarPath,
	}

	// Make the API call
	localVarHTTPResponse, err := c.options.Call(req)
	if err != nil {
		return false
	}

	// Bind the response body
	err = c.options.BindResponse(&localVarReturnValue, localVarHTTPResponse)
	if err != nil {
		return false
	}

	return true
}
