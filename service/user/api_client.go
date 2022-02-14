package user

import (
	"github.com/hasusuf/sftpgo-api/common"
	"github.com/hasusuf/sftpgo-api/types"
)

type Client interface {
	GetUsers(request types.GetUsersRequest) ([]types.User, error)
	GetUser(username string) (*types.User, error)
	CreateUser(request *types.User) (*types.User, error)
	UpdateUser(username string, request *types.User) (*types.User, error)
}

// client provides the API client for the user APIs
type client struct {
	options *common.Options
}

func New(options *common.Options) (*client, error) {
	_, err := common.ResolveAccessToken(options)
	if err != nil {
		return nil, err
	}

	common.ResolveHTTPClient(options)

	c := client{
		options: options,
	}

	return &c, nil
}
