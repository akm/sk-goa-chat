package auth

import (
	orig "firebase.google.com/go/v4/auth"
)

type (
	UserRecord = orig.UserRecord

	UserToCreate = orig.UserToCreate
	UserToUpdate = orig.UserToUpdate

	OrigClient = orig.Client
)
