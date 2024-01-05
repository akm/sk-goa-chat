//go:build !factorysetter

package auth

import (
	"fmt"
)

func SetClientFactory(f ClientFactoryFunc) {
	panic(fmt.Errorf("factorysetter is not enabled"))
}
