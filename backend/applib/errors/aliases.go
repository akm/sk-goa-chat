package errors

import (
	std "errors"

	orig "github.com/friendsofgo/errors"
)

var (
	Join = std.Join

	// As        = orig.As
	Cause  = orig.Cause
	Is     = orig.Is
	Errorf = orig.Errorf
	// WithStack = orig.WithStack
	// Wrap      = orig.Wrap
	Wrapf = orig.Wrapf
)
