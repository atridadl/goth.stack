package middleware

import (
	"net/http"

	"github.com/unrolled/secure"
	"github.com/uptrace/bunrouter"
)

func SecureHeaders(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
	secureMiddleware := secure.New(secure.Options{
		FrameDeny:          true,
		ContentTypeNosniff: true,
		BrowserXssFilter:   true,
	})

	return func(w http.ResponseWriter, req bunrouter.Request) error {
		secureMiddleware.HandlerFuncWithNext(w, req.Request, nil)
		return next(w, req)
	}
}
