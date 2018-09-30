package res

import (
	"net/http"

	"github.com/go-chi/render"
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
	render.JSON(w, r, e)
	return nil
}
