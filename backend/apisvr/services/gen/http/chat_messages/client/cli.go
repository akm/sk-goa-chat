// Code generated by goa v3.14.1, DO NOT EDIT.
//
// chat_messages HTTP client CLI support package
//
// Command:
// $ goa gen apisvr/design -o ./services

package client

import (
	chatmessages "apisvr/services/gen/chat_messages"
	"encoding/json"
	"fmt"
	"strconv"
)

// BuildListPayload builds the payload for the chat_messages list endpoint from
// CLI flags.
func BuildListPayload(chatMessagesListLimit string, chatMessagesListChannelID string, chatMessagesListAfter string, chatMessagesListBefore string, chatMessagesListIDToken string) (*chatmessages.ListPayload, error) {
	var err error
	var limit int
	{
		var v int64
		v, err = strconv.ParseInt(chatMessagesListLimit, 10, strconv.IntSize)
		limit = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for limit, must be INT")
		}
	}
	var channelID *uint64
	{
		if chatMessagesListChannelID != "" {
			val, err := strconv.ParseUint(chatMessagesListChannelID, 10, 64)
			channelID = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for channelID, must be UINT64")
			}
		}
	}
	var after *uint64
	{
		if chatMessagesListAfter != "" {
			val, err := strconv.ParseUint(chatMessagesListAfter, 10, 64)
			after = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for after, must be UINT64")
			}
		}
	}
	var before *uint64
	{
		if chatMessagesListBefore != "" {
			val, err := strconv.ParseUint(chatMessagesListBefore, 10, 64)
			before = &val
			if err != nil {
				return nil, fmt.Errorf("invalid value for before, must be UINT64")
			}
		}
	}
	var idToken string
	{
		idToken = chatMessagesListIDToken
	}
	v := &chatmessages.ListPayload{}
	v.Limit = limit
	v.ChannelID = channelID
	v.After = after
	v.Before = before
	v.IDToken = idToken

	return v, nil
}

// BuildShowPayload builds the payload for the chat_messages show endpoint from
// CLI flags.
func BuildShowPayload(chatMessagesShowID string, chatMessagesShowIDToken string) (*chatmessages.ShowPayload, error) {
	var err error
	var id uint64
	{
		id, err = strconv.ParseUint(chatMessagesShowID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT64")
		}
	}
	var idToken string
	{
		idToken = chatMessagesShowIDToken
	}
	v := &chatmessages.ShowPayload{}
	v.ID = id
	v.IDToken = idToken

	return v, nil
}

// BuildCreatePayload builds the payload for the chat_messages create endpoint
// from CLI flags.
func BuildCreatePayload(chatMessagesCreateBody string, chatMessagesCreateIDToken string) (*chatmessages.ChatMessageCreatePayload, error) {
	var err error
	var body CreateRequestBody
	{
		err = json.Unmarshal([]byte(chatMessagesCreateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"channel_id\": 3412152620045203249,\n      \"content\": \"Non accusantium.\"\n   }'")
		}
	}
	var idToken string
	{
		idToken = chatMessagesCreateIDToken
	}
	v := &chatmessages.ChatMessageCreatePayload{
		ChannelID: body.ChannelID,
		Content:   body.Content,
	}
	v.IDToken = idToken

	return v, nil
}

// BuildUpdatePayload builds the payload for the chat_messages update endpoint
// from CLI flags.
func BuildUpdatePayload(chatMessagesUpdateBody string, chatMessagesUpdateID string, chatMessagesUpdateIDToken string) (*chatmessages.ChatMessageUpdatePayload, error) {
	var err error
	var body UpdateRequestBody
	{
		err = json.Unmarshal([]byte(chatMessagesUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"content\": \"Qui similique dolor.\"\n   }'")
		}
	}
	var id uint64
	{
		id, err = strconv.ParseUint(chatMessagesUpdateID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT64")
		}
	}
	var idToken string
	{
		idToken = chatMessagesUpdateIDToken
	}
	v := &chatmessages.ChatMessageUpdatePayload{
		Content: body.Content,
	}
	v.ID = id
	v.IDToken = idToken

	return v, nil
}

// BuildDeletePayload builds the payload for the chat_messages delete endpoint
// from CLI flags.
func BuildDeletePayload(chatMessagesDeleteID string, chatMessagesDeleteIDToken string) (*chatmessages.DeletePayload, error) {
	var err error
	var id uint64
	{
		id, err = strconv.ParseUint(chatMessagesDeleteID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT64")
		}
	}
	var idToken string
	{
		idToken = chatMessagesDeleteIDToken
	}
	v := &chatmessages.DeletePayload{}
	v.ID = id
	v.IDToken = idToken

	return v, nil
}
