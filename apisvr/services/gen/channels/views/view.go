// Code generated by goa v3.14.1, DO NOT EDIT.
//
// channels views
//
// Command:
// $ goa gen apisvr/design -o ./services

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// ChannelList is the viewed result type that is projected based on a view.
type ChannelList struct {
	// Type to project
	Projected *ChannelListView
	// View to render
	View string
}

// Channel is the viewed result type that is projected based on a view.
type Channel struct {
	// Type to project
	Projected *ChannelView
	// View to render
	View string
}

// ChannelListView is a type that runs validations on a projected type.
type ChannelListView struct {
	// Items
	Items ChannelListItemCollectionView
	// Total number of items
	Total *uint64
	// Offset
	Offset *uint64
}

// ChannelListItemCollectionView is a type that runs validations on a projected
// type.
type ChannelListItemCollectionView []*ChannelListItemView

// ChannelListItemView is a type that runs validations on a projected type.
type ChannelListItemView struct {
	// ID
	ID *uint64
	// CreatedAt
	CreatedAt *string
	// UpdatedAt
	UpdatedAt *string
	// Name
	Name *string
}

// ChannelView is a type that runs validations on a projected type.
type ChannelView struct {
	// ID
	ID *uint64
	// CreatedAt
	CreatedAt *string
	// UpdatedAt
	UpdatedAt *string
	// Name
	Name *string
}

var (
	// ChannelListMap is a map indexing the attribute names of ChannelList by view
	// name.
	ChannelListMap = map[string][]string{
		"default": {
			"items",
			"total",
			"offset",
		},
	}
	// ChannelMap is a map indexing the attribute names of Channel by view name.
	ChannelMap = map[string][]string{
		"default": {
			"id",
			"created_at",
			"updated_at",
			"name",
		},
	}
	// ChannelListItemCollectionMap is a map indexing the attribute names of
	// ChannelListItemCollection by view name.
	ChannelListItemCollectionMap = map[string][]string{
		"default": {
			"id",
			"created_at",
			"updated_at",
			"name",
		},
	}
	// ChannelListItemMap is a map indexing the attribute names of ChannelListItem
	// by view name.
	ChannelListItemMap = map[string][]string{
		"default": {
			"id",
			"created_at",
			"updated_at",
			"name",
		},
	}
)

// ValidateChannelList runs the validations defined on the viewed result type
// ChannelList.
func ValidateChannelList(result *ChannelList) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateChannelListView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default"})
	}
	return
}

// ValidateChannel runs the validations defined on the viewed result type
// Channel.
func ValidateChannel(result *Channel) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateChannelView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default"})
	}
	return
}

// ValidateChannelListView runs the validations defined on ChannelListView
// using the "default" view.
func ValidateChannelListView(result *ChannelListView) (err error) {
	if result.Total == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("total", "result"))
	}
	if result.Offset == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("offset", "result"))
	}
	if result.Items != nil {
		if err2 := ValidateChannelListItemCollectionView(result.Items); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateChannelListItemCollectionView runs the validations defined on
// ChannelListItemCollectionView using the "default" view.
func ValidateChannelListItemCollectionView(result ChannelListItemCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateChannelListItemView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateChannelListItemView runs the validations defined on
// ChannelListItemView using the "default" view.
func ValidateChannelListItemView(result *ChannelListItemView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.CreatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("created_at", "result"))
	}
	if result.UpdatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("updated_at", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.CreatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.created_at", *result.CreatedAt, goa.FormatDateTime))
	}
	if result.UpdatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.updated_at", *result.UpdatedAt, goa.FormatDateTime))
	}
	return
}

// ValidateChannelView runs the validations defined on ChannelView using the
// "default" view.
func ValidateChannelView(result *ChannelView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.CreatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("created_at", "result"))
	}
	if result.UpdatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("updated_at", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.CreatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.created_at", *result.CreatedAt, goa.FormatDateTime))
	}
	if result.UpdatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.updated_at", *result.UpdatedAt, goa.FormatDateTime))
	}
	return
}
