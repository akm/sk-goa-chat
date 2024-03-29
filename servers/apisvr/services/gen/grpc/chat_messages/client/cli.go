// Code generated by goa v3.14.1, DO NOT EDIT.
//
// chat_messages gRPC client CLI support package
//
// Command:
// $ goa gen apisvr/design -o ./services

package client

import (
	chatmessages "apisvr/services/gen/chat_messages"
	chat_messagespb "apisvr/services/gen/grpc/chat_messages/pb"
	"encoding/json"
	"fmt"
)

// BuildListPayload builds the payload for the chat_messages list endpoint from
// CLI flags.
func BuildListPayload(chatMessagesListMessage string) (*chatmessages.ListPayload, error) {
	var err error
	var message chat_messagespb.ListRequest
	{
		if chatMessagesListMessage != "" {
			err = json.Unmarshal([]byte(chatMessagesListMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"after\": 178502275960762926,\n      \"before\": 13043671830441135493,\n      \"channel_id\": 1428451127207239674,\n      \"id_token\": \"abcdef12345\",\n      \"limit\": 4391472896458759808\n   }'")
			}
		}
	}
	v := &chatmessages.ListPayload{
		IDToken:   message.IdToken,
		Limit:     int(message.Limit),
		ChannelID: message.ChannelId,
		After:     message.After,
		Before:    message.Before,
	}

	return v, nil
}

// BuildShowPayload builds the payload for the chat_messages show endpoint from
// CLI flags.
func BuildShowPayload(chatMessagesShowMessage string) (*chatmessages.ShowPayload, error) {
	var err error
	var message chat_messagespb.ShowRequest
	{
		if chatMessagesShowMessage != "" {
			err = json.Unmarshal([]byte(chatMessagesShowMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"id\": 5488054146337984253,\n      \"id_token\": \"abcdef12345\"\n   }'")
			}
		}
	}
	v := &chatmessages.ShowPayload{
		IDToken: message.IdToken,
		ID:      message.Id,
	}

	return v, nil
}

// BuildCreatePayload builds the payload for the chat_messages create endpoint
// from CLI flags.
func BuildCreatePayload(chatMessagesCreateMessage string) (*chatmessages.ChatMessageCreatePayload, error) {
	var err error
	var message chat_messagespb.CreateRequest
	{
		if chatMessagesCreateMessage != "" {
			err = json.Unmarshal([]byte(chatMessagesCreateMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"channel_id\": 4983432453251475648,\n      \"content\": \"Doloremque vitae quo nesciunt necessitatibus hic.\",\n      \"id_token\": \"abcdef12345\"\n   }'")
			}
		}
	}
	v := &chatmessages.ChatMessageCreatePayload{
		IDToken:   message.IdToken,
		ChannelID: message.ChannelId,
		Content:   message.Content,
	}

	return v, nil
}

// BuildUpdatePayload builds the payload for the chat_messages update endpoint
// from CLI flags.
func BuildUpdatePayload(chatMessagesUpdateMessage string) (*chatmessages.ChatMessageUpdatePayload, error) {
	var err error
	var message chat_messagespb.UpdateRequest
	{
		if chatMessagesUpdateMessage != "" {
			err = json.Unmarshal([]byte(chatMessagesUpdateMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"content\": \"Qui voluptatibus.\",\n      \"id\": 12959863651503537545,\n      \"id_token\": \"abcdef12345\"\n   }'")
			}
		}
	}
	v := &chatmessages.ChatMessageUpdatePayload{
		IDToken: message.IdToken,
		ID:      message.Id,
		Content: message.Content,
	}

	return v, nil
}

// BuildDeletePayload builds the payload for the chat_messages delete endpoint
// from CLI flags.
func BuildDeletePayload(chatMessagesDeleteMessage string) (*chatmessages.DeletePayload, error) {
	var err error
	var message chat_messagespb.DeleteRequest
	{
		if chatMessagesDeleteMessage != "" {
			err = json.Unmarshal([]byte(chatMessagesDeleteMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"id\": 6373914294829582548,\n      \"id_token\": \"abcdef12345\"\n   }'")
			}
		}
	}
	v := &chatmessages.DeletePayload{
		IDToken: message.IdToken,
		ID:      message.Id,
	}

	return v, nil
}
