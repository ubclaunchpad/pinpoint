package res

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type baseResponse struct {
	// Basic metadata
	HTTPStatusCode int    `json:"code"`
	RequestID      string `json:"request_id,omitempty"`

	// Message is included in all responses, and is a summary of the server's response
	Message string `json:"message"`

	// Err contains additional context in the event of an error
	Err string `json:"error,omitempty"`

	// Data contains information the server wants to return
	Data interface{} `json:"data,omitempty"`
}

func newBaseResponse(
	message string,
	code int,
	kvs []interface{},
) *baseResponse {
	var data = make(map[string]interface{})
	var e string
	for i := 0; i < len(kvs)-1; i += 2 {
		var (
			k = kvs[i].(string)
			v = kvs[i+1]
		)
		if k == "error" {
			switch err := v.(type) {
			case error:
				e = err.Error()
			case string:
				e = err
			}
		} else {
			data[k] = v
		}
	}
	return &baseResponse{
		HTTPStatusCode: code,
		Message:        message,
		Err:            e,
		Data:           data,
	}
}

func (b *baseResponse) Render(w http.ResponseWriter, r *http.Request) error {
	b.RequestID = reqID(r)
	render.Status(r, b.HTTPStatusCode)
	return nil
}

func reqID(r *http.Request) string {
	if r == nil {
		return ""
	}
	return middleware.GetReqID(r.Context())
}
