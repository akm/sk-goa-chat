package cors

import (
	"net/http"
	"os"
	"strings"

	"apisvr/lib/log"
)

type Cors struct {
	allowOrigins     []string
	AllowHeaders     []string
	AllowMethods     []string
	AllowCredentials string
	Logger           *log.Logger
}

func New(allowOrigins []string) *Cors {
	return &Cors{
		allowOrigins:     allowOrigins,
		AllowHeaders:     []string{"Content-Type"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: "false",
	}
}

// NewFromEnv は環境変数から CORS を生成する
// allowOriginsEnv: 許可する Origin の環境変数名
func NewFromEnv(allowOriginsEnv string) *Cors {
	var allowOrigins []string
	if v := os.Getenv(allowOriginsEnv); v != "" {
		allowOrigins = strings.Split(v, ",")
	} else {
		allowOrigins = []string{}
	}
	return New(allowOrigins)
}

func (c *Cors) Tap(f func(*Cors)) *Cors {
	f(c)
	return c
}

func (c *Cors) WithLogger(logger *log.Logger) *Cors {
	c.Logger = logger
	return c
}

func (c *Cors) Handle(h http.Handler) http.Handler {
	if len(c.allowOrigins) == 0 {
		return h
	}

	originSet := map[string]struct{}{}
	for _, origin := range c.allowOrigins {
		originSet[origin] = struct{}{}
	}

	allowHeaders := strings.Join(c.AllowHeaders, ", ")
	allowMethods := strings.Join(c.AllowMethods, ", ")

	if c.Logger != nil {
		c.Logger.Debug().Msgf("CORS Start Handling: %+v\n", *c)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if c.Logger != nil {
			c.Logger.Debug().Msgf("CORS Header Origin: %s\n", origin)
		}
		if origin != "" {
			if _, ok := originSet[origin]; ok {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Headers", allowHeaders)
				w.Header().Set("Access-Control-Allow-Methods", allowMethods)
				w.Header().Set("Access-Control-Allow-Credentials", c.AllowCredentials)
			}
			// Preflight Request の場合は 200 を返す
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				w.WriteHeader(http.StatusOK)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}
