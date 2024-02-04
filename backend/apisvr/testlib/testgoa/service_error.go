package testgoa

import (
	"apisvr/applib/errors"
	"testing"

	"github.com/stretchr/testify/assert"
	goa "goa.design/goa/v3/pkg"
)

func AssertServiceError(t *testing.T, expectedName string, err error) {
	if assert.Error(t, err) {
		cause := errors.Cause(err)
		if assert.IsType(t, (*goa.ServiceError)(nil), cause) {
			assert.Equal(t, expectedName, cause.(*goa.ServiceError).Name)
		}
	}
}
