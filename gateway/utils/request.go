package utils

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
)

// RequestID fetches request-id set by chi's RID middleware
func RequestID(r *http.Request) string {
	if reqID := r.Context().Value(middleware.RequestIDKey); reqID != nil {
		return reqID.(string)
	}
	return ""
}
