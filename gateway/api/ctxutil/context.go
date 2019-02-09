package ctxutil

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
)

// Key is a useful type for denoting context keys
type Key string

func GetRequestID(r *http.Request) (requestID string) {
	if reqID := r.Context().Value(middleware.RequestIDKey); reqID != nil {
		requestID = reqID.(string)
	}
	return
}
