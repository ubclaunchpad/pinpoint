package res

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/ubclaunchpad/pinpoint/gateway/utils"
)

// MsgResponse is the template for a typical HTTP response for messages
type MsgResponse struct {
	Message        string `json:"message"`
	HTTPStatusCode int    `json:"-"`
	RequestID      string `json:"request-id,omitempty"`

	Details map[string]interface{} `json:"details,omitempty"`
}

// Render renders a MsgResponse
func (m *MsgResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, m.HTTPStatusCode)
	return nil
}

// Message is a shortcut for non-error statuses
func Message(r *http.Request, msg string, code int, fields ...interface{}) render.Renderer {
	return &MsgResponse{
		HTTPStatusCode: code,
		Message:        msg,
		RequestID:      utils.RequestID(r),
		Details:        utils.ToMap(fields),
	}
}
