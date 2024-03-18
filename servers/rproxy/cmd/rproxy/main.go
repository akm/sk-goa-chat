package main

import (
	"applib/collection"
	"context"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"applib/errors"
	"applib/firebase"
	"applib/firebase/auth"
	"applib/log"
)

func main() {
	hostPort := os.Getenv("HOST_PORT")
	if hostPort == "" {
		hostPort = ":8080"
	}

	tokenHeaderKey := os.Getenv("TOKEN_HEADER_KEY")
	if tokenHeaderKey == "" {
		tokenHeaderKey = "X-ID-TOKEN"
	}

	uidHeaderKey := os.Getenv("UID_HEADER_KEY")
	if uidHeaderKey == "" {
		uidHeaderKey = "X-UID"
	}

	logger := log.New(os.Stderr, "rproxy: ", log.LstdFlags)

	// https://gist.github.com/JalfResi/6287706
	// https://pkg.go.dev/net/http/httputil#ReverseProxy

	verifyIdToken := verifyIdTokenFunc(logger, tokenHeaderKey)
	handler := func(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.URL)
			if !skipAuth(r.URL) {
				ctx := r.Context()
				uid, err := verifyIdToken(ctx, r)
				if err != nil {
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
				}
				r.Header.Set(uidHeaderKey, uid)
			}
			p.ServeHTTP(w, r)
		}
	}

	proxy := newMultiHostReverseProxy()
	http.HandleFunc("/", handler(proxy))
	err := http.ListenAndServe(hostPort, nil)
	if err != nil {
		panic(err)
	}
}

var noAuthPaths = []string{
	"/signin",
	"/signup",
	"/api/version",
}

var (
	uisvrOriginUrl  = GetUrlFromEnv("UISVR_ORIGIN_URL")
	apisvrOriginUrl = GetUrlFromEnv("APISVR_ORIGIN_URL")
)

func GetUrlFromEnv(key string) *url.URL {
	s := os.Getenv(key)
	r, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	return r
}

func mappingUrl(u *url.URL) *url.URL {
	if strings.HasPrefix(u.Path, "/api/") {
		return apisvrOriginUrl
	}
	return uisvrOriginUrl
}

func skipAuth(u *url.URL) bool {
	return collection.Any(noAuthPaths, func(p string) bool { return u.Path == p })
}

func newMultiHostReverseProxy() *httputil.ReverseProxy {
	director := func(req *http.Request) {
		log.Println(req.URL)
		targetUrl := mappingUrl(req.URL)
		rewriteRequestURL(req, targetUrl)
	}
	return &httputil.ReverseProxy{Director: director}
}

var authError = errors.New("auth error")

func verifyIdTokenFunc(logger *log.Logger, idTokenHeader string) func(ctx context.Context, r *http.Request) (string, error) {
	return func(ctx context.Context, r *http.Request) (string, error) {
		fbapp, err := firebase.NewApp(ctx, nil)
		if err != nil {
			return "", errors.Wrapf(err, "firebase.NewApp")
		}
		fbauth, err := auth.NewClientWithLogger(ctx, fbapp, logger)
		if err != nil {
			return "", errors.Wrapf(err, "auth.NewClientWithLogger")
		}

		givenIdToken := r.Header.Get(idTokenHeader)
		token, err := fbauth.VerifyIDToken(ctx, givenIdToken)
		if err != nil {
			return "", authError
		}
		return token.UID, nil
	}
}

// 以下 src/net/http/httputil/reverseproxy.go からコピー

func rewriteRequestURL(req *http.Request, target *url.URL) {
	targetQuery := target.RawQuery
	req.URL.Scheme = target.Scheme
	req.URL.Host = target.Host
	req.URL.Path, req.URL.RawPath = joinURLPath(target, req.URL)
	if targetQuery == "" || req.URL.RawQuery == "" {
		req.URL.RawQuery = targetQuery + req.URL.RawQuery
	} else {
		req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
	}
}

func joinURLPath(a, b *url.URL) (path, rawpath string) {
	if a.RawPath == "" && b.RawPath == "" {
		return singleJoiningSlash(a.Path, b.Path), ""
	}
	// Same as singleJoiningSlash, but uses EscapedPath to determine
	// whether a slash should be added
	apath := a.EscapedPath()
	bpath := b.EscapedPath()

	aslash := strings.HasSuffix(apath, "/")
	bslash := strings.HasPrefix(bpath, "/")

	switch {
	case aslash && bslash:
		return a.Path + b.Path[1:], apath + bpath[1:]
	case !aslash && !bslash:
		return a.Path + "/" + b.Path, apath + "/" + bpath
	}
	return a.Path + b.Path, apath + bpath
}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}
