// Code generated by goa v3.14.1, DO NOT EDIT.
//
// channels gRPC client CLI support package
//
// Command:
// $ goa gen apisvr/design -o ./services

package client

import (
	channels "apisvr/services/gen/channels"
	channelspb "apisvr/services/gen/grpc/channels/pb"
	"encoding/json"
	"fmt"
)

// BuildListPayload builds the payload for the channels list endpoint from CLI
// flags.
func BuildListPayload(channelsListMessage string) (*channels.ListPayload, error) {
	var err error
	var message channelspb.ListRequest
	{
		if channelsListMessage != "" {
			err = json.Unmarshal([]byte(channelsListMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"session_id\": \"Alias rem.\"\n   }'")
			}
		}
	}
	v := &channels.ListPayload{
		SessionID: message.SessionId,
	}

	return v, nil
}

// BuildShowPayload builds the payload for the channels show endpoint from CLI
// flags.
func BuildShowPayload(channelsShowMessage string) (*channels.ShowPayload, error) {
	var err error
	var message channelspb.ShowRequest
	{
		if channelsShowMessage != "" {
			err = json.Unmarshal([]byte(channelsShowMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"id\": 13315485509768610097,\n      \"session_id\": \"Tempore et minus molestias aspernatur blanditiis neque.\"\n   }'")
			}
		}
	}
	v := &channels.ShowPayload{
		SessionID: message.SessionId,
		ID:        message.Id,
	}

	return v, nil
}

// BuildCreatePayload builds the payload for the channels create endpoint from
// CLI flags.
func BuildCreatePayload(channelsCreateMessage string) (*channels.ChannelCreatePayload, error) {
	var err error
	var message channelspb.CreateRequest
	{
		if channelsCreateMessage != "" {
			err = json.Unmarshal([]byte(channelsCreateMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"name\": \"Consequatur consequatur doloremque assumenda sint consequuntur corrupti.\",\n      \"session_id\": \"Blanditiis molestias.\"\n   }'")
			}
		}
	}
	v := &channels.ChannelCreatePayload{
		SessionID: message.SessionId,
		Name:      message.Name,
	}

	return v, nil
}

// BuildUpdatePayload builds the payload for the channels update endpoint from
// CLI flags.
func BuildUpdatePayload(channelsUpdateMessage string) (*channels.ChannelUpdatePayload, error) {
	var err error
	var message channelspb.UpdateRequest
	{
		if channelsUpdateMessage != "" {
			err = json.Unmarshal([]byte(channelsUpdateMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"id\": 5096907298334997690,\n      \"name\": \"Commodi et voluptas.\",\n      \"session_id\": \"Cumque magni vitae.\"\n   }'")
			}
		}
	}
	v := &channels.ChannelUpdatePayload{
		SessionID: message.SessionId,
		ID:        message.Id,
		Name:      message.Name,
	}

	return v, nil
}

// BuildDeletePayload builds the payload for the channels delete endpoint from
// CLI flags.
func BuildDeletePayload(channelsDeleteMessage string) (*channels.DeletePayload, error) {
	var err error
	var message channelspb.DeleteRequest
	{
		if channelsDeleteMessage != "" {
			err = json.Unmarshal([]byte(channelsDeleteMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"id\": 10482655633392866197,\n      \"session_id\": \"Fugiat laborum expedita et veritatis itaque.\"\n   }'")
			}
		}
	}
	v := &channels.DeletePayload{
		SessionID: message.SessionId,
		ID:        message.Id,
	}

	return v, nil
}
