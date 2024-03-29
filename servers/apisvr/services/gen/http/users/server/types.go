// Code generated by goa v3.14.1, DO NOT EDIT.
//
// users HTTP server types
//
// Command:
// $ goa gen apisvr/design -o ./services

package server

import (
	users "apisvr/services/gen/users"
	usersviews "apisvr/services/gen/users/views"

	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "users" service "create" endpoint HTTP
// request body.
type CreateRequestBody struct {
	// Name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Email
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
}

// ListResponseBody is the type of the "users" service "list" endpoint HTTP
// response body.
type ListResponseBody struct {
	// Items
	Items UserListItemResponseBodyCollection `form:"items" json:"items" xml:"items"`
	// Total number of items
	Total uint64 `form:"total" json:"total" xml:"total"`
	// Offset
	Offset uint64 `form:"offset" json:"offset" xml:"offset"`
}

// CreateResponseBody is the type of the "users" service "create" endpoint HTTP
// response body.
type CreateResponseBody struct {
	// ID
	ID uint64 `form:"id" json:"id" xml:"id"`
	// CreatedAt
	CreatedAt string `form:"created_at" json:"created_at" xml:"created_at"`
	// UpdatedAt
	UpdatedAt string `form:"updated_at" json:"updated_at" xml:"updated_at"`
	// Name
	Name string `form:"name" json:"name" xml:"name"`
}

// CreateInvalidPayloadResponseBody is the type of the "users" service "create"
// endpoint HTTP response body for the "invalid_payload" error.
type CreateInvalidPayloadResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UserListItemResponseBodyCollection is used to define fields on response body
// types.
type UserListItemResponseBodyCollection []*UserListItemResponseBody

// UserListItemResponseBody is used to define fields on response body types.
type UserListItemResponseBody struct {
	// ID
	ID uint64 `form:"id" json:"id" xml:"id"`
	// Name
	Name string `form:"name" json:"name" xml:"name"`
}

// NewListResponseBody builds the HTTP response body from the result of the
// "list" endpoint of the "users" service.
func NewListResponseBody(res *usersviews.UserListView) *ListResponseBody {
	body := &ListResponseBody{
		Total:  *res.Total,
		Offset: *res.Offset,
	}
	if res.Items != nil {
		body.Items = make([]*UserListItemResponseBody, len(res.Items))
		for i, val := range res.Items {
			body.Items[i] = marshalUsersviewsUserListItemViewToUserListItemResponseBody(val)
		}
	} else {
		body.Items = []*UserListItemResponseBody{}
	}
	return body
}

// NewCreateResponseBody builds the HTTP response body from the result of the
// "create" endpoint of the "users" service.
func NewCreateResponseBody(res *usersviews.UserView) *CreateResponseBody {
	body := &CreateResponseBody{
		ID:        *res.ID,
		CreatedAt: *res.CreatedAt,
		UpdatedAt: *res.UpdatedAt,
		Name:      *res.Name,
	}
	return body
}

// NewCreateInvalidPayloadResponseBody builds the HTTP response body from the
// result of the "create" endpoint of the "users" service.
func NewCreateInvalidPayloadResponseBody(res *goa.ServiceError) *CreateInvalidPayloadResponseBody {
	body := &CreateInvalidPayloadResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateUserCreatePayload builds a users service create endpoint payload.
func NewCreateUserCreatePayload(body *CreateRequestBody) *users.UserCreatePayload {
	v := &users.UserCreatePayload{
		Name:  *body.Name,
		Email: *body.Email,
	}

	return v
}

// ValidateCreateRequestBody runs the validations defined on CreateRequestBody
func ValidateCreateRequestBody(body *CreateRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	return
}
