package res

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/ubclaunchpad/pinpoint/gateway/utils"
)

// Err is a basic error response constructor
func Err(r *http.Request, err error, status int, msg ...interface{}) render.Renderer {
	return &ErrResponse{
		HTTPStatusCode: status,
		StatusText:     utils.FirstString(msg),
		ErrorText:      err.Error(),
		RequestID:      utils.RequestID(r),
	}
}

// ErrInternalServer is a shortcut for internal server errors
func ErrInternalServer(r *http.Request, err error, msg ...interface{}) render.Renderer {
	return &ErrResponse{
		HTTPStatusCode: http.StatusInternalServerError,
		StatusText:     utils.FirstString(msg),
		ErrorText:      err.Error(),
		RequestID:      utils.RequestID(r),
	}
}

// ErrBadRequest is a shortcut for internal server errors
func ErrBadRequest(r *http.Request, err error, msg string, missingFields ...string) render.Renderer {
	return &ErrResponse{
		HTTPStatusCode: http.StatusBadRequest,
		StatusText:     msg,
		ErrorText:      err.Error(),
		RequestID:      utils.RequestID(r),
	}
}

// Message is a shortcut for non-error statuses
func Message(r *http.Request, msg string, code int, fields ...interface{}) render.Renderer {
	return &MsgResponse{
		HTTPStatusCode: code,
		StatusText:     msg,
		RequestID:      utils.RequestID(r),
		Details:        utils.ToMap(fields),
	}
}
