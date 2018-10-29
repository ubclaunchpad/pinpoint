package res

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/ubclaunchpad/pinpoint/gateway/utils"
)

// Err is a basic error response constructor
func Err(r *http.Request, err error, status int, msg ...string) render.Renderer {
	return &ErrResponse{
		HTTPStatusCode: status,
		StatusText:     utils.FirstString(msg),
		ErrorText:      err.Error(),
		RequestID:      utils.RequestID(r),
	}
}

// ErrInternalServer is a shortcut for internal server errors
func ErrInternalServer(r *http.Request, err error, msg ...string) render.Renderer {
	return &ErrResponse{
		HTTPStatusCode: http.StatusInternalServerError,
		StatusText:     utils.FirstString(msg),
		ErrorText:      err.Error(),
		RequestID:      utils.RequestID(r),
	}
}

func Message(r *http.Request, msg ...string) render.Renderer {
	return &MsgResponse{
		HTTPStatusCode: http.StatusCreated,
		StatusText:     utils.FirstString(msg),
		RequestID:      utils.RequestID(r),
	}
}
