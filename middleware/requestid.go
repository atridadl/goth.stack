package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/uptrace/bunrouter"
)

type contextKey string

func (c contextKey) String() string {
	return string(c)
}

var (
	HeaderXRequestID = "X-Request-ID"
	requestIDKey     = contextKey("requestID")
)

func RequestID(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		reqID := req.Header.Get(HeaderXRequestID)
		if reqID == "" {
			reqID = uuid.New().String()
		}

		ctx := context.WithValue(req.Context(), requestIDKey, reqID)

		w.Header().Set(HeaderXRequestID, reqID)
		return next(w, req.WithContext(ctx))
	}
}

func GetRequestID(ctx context.Context) string {
	if reqID, ok := ctx.Value(requestIDKey).(string); ok {
		return reqID
	}
	return ""
}
