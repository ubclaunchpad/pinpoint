package res

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// BaseResponse is the underlying structure of all API responses
type BaseResponse struct {
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
) *BaseResponse {
	e, data := formatData(kvs)
	return &BaseResponse{
		HTTPStatusCode: code,
		Message:        message,
		Err:            e,
		Data:           data,
	}
}

// Render implements chi/render.Render
func (b *BaseResponse) Render(w http.ResponseWriter, r *http.Request) error {
	b.RequestID = reqID(r)
	render.Status(r, b.HTTPStatusCode)
	return nil
}

func formatData(kvs []interface{}) (e string, d interface{}) {
	if len(kvs) < 1 {
		return "", nil
	}

	var data = make(map[string]interface{})
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

	if len(data) < 1 {
		return e, nil
	}
	return e, data
}

func reqID(r *http.Request) string {
	if r == nil || r.Context() == nil {
		return ""
	}
	return middleware.GetReqID(r.Context())
}
