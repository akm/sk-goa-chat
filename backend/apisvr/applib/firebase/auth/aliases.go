package auth

import (
	orig "firebase.google.com/go/v4/auth"
)

type (
	OrigClient = orig.Client

	UserRecord = orig.UserRecord

	UserToCreate = orig.UserToCreate
	UserToUpdate = orig.UserToUpdate

	UserIterator       = orig.UserIterator
	ExportedUserRecord = orig.ExportedUserRecord

	Token = orig.Token
)
