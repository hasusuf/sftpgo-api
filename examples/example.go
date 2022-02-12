package main

import (
	"encoding/base64"
	"fmt"

	"github.com/hasusuf/sftpgo-api/service/user"
	"github.com/hasusuf/sftpgo-api/types"
)

const (
	host     = "sftpgo.example.com:8080"
	username = "admin"
	password = "PLAIN-TEXT-PASSWORD-NOT-SECURE"
)

func ExampleCreateUser() {
	// Instantiate client
	cli, err := user.New(host, username, password, user.HTTP)
	if err != nil {
		panic(err)
	}

	createUserPayload := `
{
  "status": 1,
  "username": "foo",
  "email": "foo@example.com",
  "expiration_date": 0,
  "password": "SECRET-PASSWORD",
  "public_keys": [],
  "virtual_folders": [],
  "uid": 0,
  "gid": 0,
  "max_sessions": 0,
  "quota_size": 0,
  "quota_files": 0,
  "permissions": {
    "/": [
      "*"
    ]
  },
  "used_quota_size": 0,
  "used_quota_files": 0,
  "last_quota_update": 0,
  "upload_bandwidth": 0,
  "download_bandwidth": 0,
  "filters": {
    "denied_protocols": [
      "DAV"
    ],
    "allow_api_key_auth": false
  },
  "filesystem": {
    "provider": 1,
    "s3config": {
      "bucket": "my-bucket-name",
      "region": "us-west-1",
      "access_key": "ACCESS-KEY-ID",
      "access_secret": {
        "status": "Plain",
        "payload": "ACCESS-KEY-SECRET"
      },
      "endpoint": "https://minio.example.com:9000",
      "storage_class": "Standard"
    }
  }
}
`

	newUser := types.NewUser()
	err = types.NewNullableUser(newUser).UnmarshalJSON([]byte(createUserPayload))
	if err != nil {
		panic(err)
	}

	response, err := cli.CreateUser(newUser)
	if err != nil {
		fmt.Println(err)
	} else {
		json, err := response.MarshalJSON()
		if err != nil {
			panic(err)
		}
		fmt.Println(string(json))
	}
}

func ExampleUpdateUser() {
	// Instantiate client
	cli, err := user.New(host, username, password, user.HTTP)
	if err != nil {
		panic(err)
	}

	updateUserPayload := `
{
  "status": 0,
  "description": "updated user",
  "filesystem": {
    "provider": 1,
    "s3config": {
      "bucket": "my-bucket-name",
      "region": "us-west-1",
      "access_key": "ACCESS-KEY-ID",
      "access_secret": {
        "status": "Plain",
        "payload": "ACCESS-KEY-SECRET"
      },
      "endpoint": "https://minio.example.com:9000",
      "storage_class": "Standard"
    }
  }
}
`

	user := types.NewUser()
	err = types.NewNullableUser(user).UnmarshalJSON([]byte(updateUserPayload))
	if err != nil {
		panic(err)
	}

	response, err := cli.UpdateUser("foo", user)
	if err != nil {
		fmt.Println(err)
	} else {
		json, err := response.MarshalJSON()
		if err != nil {
			panic(err)
		}
		fmt.Println(string(json))
	}
}

func ExampleGetAccessTokenWithoutClient() {
	// Build credentials
	data := []byte(fmt.Sprintf("%s:%s", username, password))
	token := base64.StdEncoding.EncodeToString(data)

	// Build the header
	Headers := make(map[string]string)
	Headers["Authorization"] = "Basic " + token

	// Build the request params
	method := user.GET
	scheme := user.HTTP
	endpoint := "/api/v2/token"

	// Make the API call
	request := user.Request{
		Method:   method,
		Host:     host,
		Scheme:   &scheme,
		Endpoint: endpoint,
		Headers:  Headers,
	}

	response, err := user.Call(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
	}
}
