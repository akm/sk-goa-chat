package errors

import (
	orig "github.com/friendsofgo/errors"
)

var (
	// As        = orig.As
	Cause  = orig.Cause
	Is     = orig.Is
	Errorf = orig.Errorf
	// WithStack = orig.WithStack
	// Wrap      = orig.Wrap
	Wrapf = orig.Wrapf
)
