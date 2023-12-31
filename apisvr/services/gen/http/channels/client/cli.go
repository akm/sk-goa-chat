// Code generated by goa v3.14.1, DO NOT EDIT.
//
// channels HTTP client CLI support package
//
// Command:
// $ goa gen apisvr/design -o ./services

package client

import (
	channels "apisvr/services/gen/channels"
	"encoding/json"
	"fmt"
	"strconv"
)

// BuildShowPayload builds the payload for the channels show endpoint from CLI
// flags.
func BuildShowPayload(channelsShowID string) (*channels.ShowPayload, error) {
	var err error
	var id uint64
	{
		id, err = strconv.ParseUint(channelsShowID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT64")
		}
	}
	v := &channels.ShowPayload{}
	v.ID = id

	return v, nil
}

// BuildCreatePayload builds the payload for the channels create endpoint from
// CLI flags.
func BuildCreatePayload(channelsCreateBody string) (*channels.ChannelCreatePayload, error) {
	var err error
	var body CreateRequestBody
	{
		err = json.Unmarshal([]byte(channelsCreateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"name\": \"Enim iste odio.\"\n   }'")
		}
	}
	v := &channels.ChannelCreatePayload{
		Name: body.Name,
	}

	return v, nil
}

// BuildUpdatePayload builds the payload for the channels update endpoint from
// CLI flags.
func BuildUpdatePayload(channelsUpdateBody string, channelsUpdateID string) (*channels.ChannelUpdatePayload, error) {
	var err error
	var body UpdateRequestBody
	{
		err = json.Unmarshal([]byte(channelsUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"name\": \"Consequatur officia dolores corporis.\"\n   }'")
		}
	}
	var id uint64
	{
		id, err = strconv.ParseUint(channelsUpdateID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT64")
		}
	}
	v := &channels.ChannelUpdatePayload{
		Name: body.Name,
	}
	v.ID = id

	return v, nil
}

// BuildDeletePayload builds the payload for the channels delete endpoint from
// CLI flags.
func BuildDeletePayload(channelsDeleteID string) (*channels.DeletePayload, error) {
	var err error
	var id uint64
	{
		id, err = strconv.ParseUint(channelsDeleteID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT64")
		}
	}
	v := &channels.DeletePayload{}
	v.ID = id

	return v, nil
}
