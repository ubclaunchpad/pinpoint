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
	return nil
}

// MsgResponse is the template for a typical HTTP response for messages
type MsgResponse struct {
	Msg            []string `json:"msg"`
	HTTPStatusCode int      `json:"-"`
	StatusText     string   `json:"status"`
	RequestID      string   `json:"request-id,omitempty"`
}

// Render renders a MsgResponse
func (m *MsgResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, m.HTTPStatusCode)
	return nil
}
