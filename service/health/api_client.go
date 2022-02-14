package health

import (
	"github.com/hasusuf/sftpgo-api/common"
)

type Client interface {
	Health() (ok bool)
}

// Client allows modification of client headers and other settings
// See https://golang.org/pkg/net/http
type client struct {
	options *common.Options
}

func New(options *common.Options) (*client, error) {
	common.ResolveHTTPClient(options)

	c := client{
		options: options,
	}

	return &c, nil
}
