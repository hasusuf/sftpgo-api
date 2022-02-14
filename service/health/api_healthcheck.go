package health

import (
	"github.com/hasusuf/sftpgo-api/common"
	"net/http"
)

func (c *client) Health() (ok bool) {
	localVarPath := "/healthz"

	req := common.Request{
		Method:   common.GET,
		Endpoint: localVarPath,
	}

	// Make the API call
	localVarHTTPResponse, err := c.options.Call(req)
	if err != nil {
		return false
	}

	return localVarHTTPResponse.StatusCode == http.StatusOK
}
