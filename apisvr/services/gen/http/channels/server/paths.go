// Code generated by goa v3.14.0, DO NOT EDIT.
//
// HTTP request path constructors for the channels service.
//
// Command:
// $ goa gen apisvr/design -o ./services

package server

import (
	"fmt"
)

// ListChannelsPath returns the URL path to the channels service list HTTP endpoint.
func ListChannelsPath() string {
	return "/api/channels"
}

// ShowChannelsPath returns the URL path to the channels service show HTTP endpoint.
func ShowChannelsPath(id uint64) string {
	return fmt.Sprintf("/api/channels/%v", id)
}

// CreateChannelsPath returns the URL path to the channels service create HTTP endpoint.
func CreateChannelsPath() string {
	return "/api/channels"
}

// UpdateChannelsPath returns the URL path to the channels service update HTTP endpoint.
func UpdateChannelsPath(id uint64) string {
	return fmt.Sprintf("/api/channels/%v", id)
}

// DeleteChannelsPath returns the URL path to the channels service delete HTTP endpoint.
func DeleteChannelsPath(id uint64) string {
	return fmt.Sprintf("/api/channels/%v", id)
}
