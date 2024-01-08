// Code generated by goa v3.14.1, DO NOT EDIT.
//
// chat_messages views
//
// Command:
// $ goa gen apisvr/design -o ./services

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// ChatMessageList is the viewed result type that is projected based on a view.
type ChatMessageList struct {
	// Type to project
	Projected *ChatMessageListView
	// View to render
	View string
}

// ChatMessage is the viewed result type that is projected based on a view.
type ChatMessage struct {
	// Type to project
	Projected *ChatMessageView
	// View to render
	View string
}

// ChatMessageListView is a type that runs validations on a projected type.
type ChatMessageListView struct {
	// Items
	Items ChatMessageListItemCollectionView
}

// ChatMessageListItemCollectionView is a type that runs validations on a
// projected type.
type ChatMessageListItemCollectionView []*ChatMessageListItemView

// ChatMessageListItemView is a type that runs validations on a projected type.
type ChatMessageListItemView struct {
	// ID
	ID *uint64
	// CreatedAt
	CreatedAt *string
	// UpdatedAt
	UpdatedAt *string
	// Channel ID
	ChannelID *uint64
	// User ID
	UserID *uint64
	// User Name
	UserName *uint64
	// Content
	Content *string
}

// ChatMessageView is a type that runs validations on a projected type.
type ChatMessageView struct {
	// ID
	ID *uint64
	// CreatedAt
	CreatedAt *string
	// UpdatedAt
	UpdatedAt *string
	// Channel ID
	ChannelID *uint64
	// User ID
	UserID *uint64
	// User Name
	UserName *uint64
	// Content
	Content *string
}

var (
	// ChatMessageListMap is a map indexing the attribute names of ChatMessageList
	// by view name.
	ChatMessageListMap = map[string][]string{
		"default": {
			"items",
		},
	}
	// ChatMessageMap is a map indexing the attribute names of ChatMessage by view
	// name.
	ChatMessageMap = map[string][]string{
		"default": {
			"id",
			"created_at",
			"updated_at",
			"channel_id",
			"user_id",
			"user_name",
			"content",
		},
	}
	// ChatMessageListItemCollectionMap is a map indexing the attribute names of
	// ChatMessageListItemCollection by view name.
	ChatMessageListItemCollectionMap = map[string][]string{
		"default": {
			"id",
			"created_at",
			"updated_at",
			"channel_id",
			"user_id",
			"user_name",
			"content",
		},
	}
	// ChatMessageListItemMap is a map indexing the attribute names of
	// ChatMessageListItem by view name.
	ChatMessageListItemMap = map[string][]string{
		"default": {
			"id",
			"created_at",
			"updated_at",
			"channel_id",
			"user_id",
			"user_name",
			"content",
		},
	}
)

// ValidateChatMessageList runs the validations defined on the viewed result
// type ChatMessageList.
func ValidateChatMessageList(result *ChatMessageList) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateChatMessageListView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default"})
	}
	return
}

// ValidateChatMessage runs the validations defined on the viewed result type
// ChatMessage.
func ValidateChatMessage(result *ChatMessage) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateChatMessageView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default"})
	}
	return
}

// ValidateChatMessageListView runs the validations defined on
// ChatMessageListView using the "default" view.
func ValidateChatMessageListView(result *ChatMessageListView) (err error) {

	if result.Items != nil {
		if err2 := ValidateChatMessageListItemCollectionView(result.Items); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateChatMessageListItemCollectionView runs the validations defined on
// ChatMessageListItemCollectionView using the "default" view.
func ValidateChatMessageListItemCollectionView(result ChatMessageListItemCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateChatMessageListItemView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateChatMessageListItemView runs the validations defined on
// ChatMessageListItemView using the "default" view.
func ValidateChatMessageListItemView(result *ChatMessageListItemView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.CreatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("created_at", "result"))
	}
	if result.UpdatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("updated_at", "result"))
	}
	if result.ChannelID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("channel_id", "result"))
	}
	if result.UserID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("user_id", "result"))
	}
	if result.UserName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("user_name", "result"))
	}
	if result.Content == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("content", "result"))
	}
	if result.CreatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.created_at", *result.CreatedAt, goa.FormatDateTime))
	}
	if result.UpdatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.updated_at", *result.UpdatedAt, goa.FormatDateTime))
	}
	return
}

// ValidateChatMessageView runs the validations defined on ChatMessageView
// using the "default" view.
func ValidateChatMessageView(result *ChatMessageView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.CreatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("created_at", "result"))
	}
	if result.UpdatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("updated_at", "result"))
	}
	if result.ChannelID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("channel_id", "result"))
	}
	if result.UserID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("user_id", "result"))
	}
	if result.UserName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("user_name", "result"))
	}
	if result.Content == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("content", "result"))
	}
	if result.CreatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.created_at", *result.CreatedAt, goa.FormatDateTime))
	}
	if result.UpdatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.updated_at", *result.UpdatedAt, goa.FormatDateTime))
	}
	return
}
