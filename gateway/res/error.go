package res

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/ubclaunchpad/pinpoint/gateway/utils"
)

// ErrResponse is the template for a typical HTTP response for errors
type ErrResponse struct {
	Err            error  `json:"-"`
	HTTPStatusCode int    `json:"-"`
	StatusText     string `json:"status"`
	ErrorText      string `json:"error,omitempty"`
	RequestID      string `json:"request-id,omitempty"`
}

// Render renders an ErrResponse
func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

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

// ErrBadRequest is a shortcut for bad requests
func ErrBadRequest(r *http.Request, err error, msg string, missingFields ...string) render.Renderer {
	return &ErrResponse{
		HTTPStatusCode: http.StatusBadRequest,
		StatusText:     msg,
		ErrorText:      err.Error(),
		RequestID:      utils.RequestID(r),
	}
}

// ErrUnauthorized is a shortcut for unauthorized requests
func ErrUnauthorized(r *http.Request, err error, msg string, missingFields ...string) render.Renderer {
	return &ErrResponse{
		HTTPStatusCode: http.StatusUnauthorized,
		StatusText:     msg,
		ErrorText:      err.Error(),
		RequestID:      utils.RequestID(r),
	}
}
