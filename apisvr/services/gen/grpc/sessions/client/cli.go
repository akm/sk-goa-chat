// Code generated by goa v3.14.1, DO NOT EDIT.
//
// sessions gRPC client CLI support package
//
// Command:
// $ goa gen apisvr/design -o ./services

package client

import (
	sessionspb "apisvr/services/gen/grpc/sessions/pb"
	sessions "apisvr/services/gen/sessions"
	"encoding/json"
	"fmt"
)

// BuildCreatePayload builds the payload for the sessions create endpoint from
// CLI flags.
func BuildCreatePayload(sessionsCreateMessage string) (*sessions.CreatePayload, error) {
	var err error
	var message sessionspb.CreateRequest
	{
		if sessionsCreateMessage != "" {
			err = json.Unmarshal([]byte(sessionsCreateMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"id_token\": \"Aliquid ea nulla fugit porro recusandae velit.\"\n   }'")
			}
		}
	}
	v := &sessions.CreatePayload{
		IDToken: message.IdToken,
	}

	return v, nil
}

// BuildDeletePayload builds the payload for the sessions delete endpoint from
// CLI flags.
func BuildDeletePayload(sessionsDeleteMessage string) (*sessions.DeletePayload, error) {
	var err error
	var message sessionspb.DeleteRequest
	{
		if sessionsDeleteMessage != "" {
			err = json.Unmarshal([]byte(sessionsDeleteMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"session_id\": \"Quas voluptatibus consequatur nemo earum vero in.\"\n   }'")
			}
		}
	}
	v := &sessions.DeletePayload{
		SessionID: message.SessionId,
	}

	return v, nil
}
