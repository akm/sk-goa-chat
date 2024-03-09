//go:build factorysetter

package auth

func SetClientFactory(f ClientFactoryFunc) {
	clientFactory = f
}
