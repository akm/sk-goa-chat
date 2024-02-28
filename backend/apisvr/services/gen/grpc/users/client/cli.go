// Code generated by goa v3.14.1, DO NOT EDIT.
//
// users gRPC client CLI support package
//
// Command:
// $ goa gen apisvr/design -o ./services

package client

import (
	userspb "apisvr/services/gen/grpc/users/pb"
	users "apisvr/services/gen/users"
	"encoding/json"
	"fmt"
)

// BuildCreatePayload builds the payload for the users create endpoint from CLI
// flags.
func BuildCreatePayload(usersCreateMessage string) (*users.UserCreatePayload, error) {
	var err error
	var message userspb.CreateRequest
	{
		if usersCreateMessage != "" {
			err = json.Unmarshal([]byte(usersCreateMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"email\": \"Delectus natus libero consequatur ut ad cupiditate.\",\n      \"name\": \"Est aperiam.\"\n   }'")
			}
		}
	}
	v := &users.UserCreatePayload{
		Name:  message.Name,
		Email: message.Email,
	}

	return v, nil
}
