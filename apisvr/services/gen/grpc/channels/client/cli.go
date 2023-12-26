// Code generated by goa v3.14.0, DO NOT EDIT.
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

// BuildShowPayload builds the payload for the channels show endpoint from CLI
// flags.
func BuildShowPayload(channelsShowMessage string) (*channels.ShowPayload, error) {
	var err error
	var message channelspb.ShowRequest
	{
		if channelsShowMessage != "" {
			err = json.Unmarshal([]byte(channelsShowMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"id\": 7179248999906013810\n   }'")
			}
		}
	}
	v := &channels.ShowPayload{
		ID: message.Id,
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
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"name\": \"Ipsum adipisci sunt qui nisi.\"\n   }'")
			}
		}
	}
	v := &channels.ChannelCreatePayload{
		Name: message.Name,
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
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"id\": 18394200013453555426,\n      \"name\": \"Qui dolor nam labore odit et reiciendis.\"\n   }'")
			}
		}
	}
	v := &channels.ChannelUpdatePayload{
		ID:   message.Id,
		Name: message.Name,
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
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"id\": 9178827906163796587\n   }'")
			}
		}
	}
	v := &channels.DeletePayload{
		ID: message.Id,
	}

	return v, nil
}
