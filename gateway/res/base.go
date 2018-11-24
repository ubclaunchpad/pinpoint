package res

import "net/http"

// BaseResponse is the underlying response structure to all responses
type BaseResponse struct {
	HTTPStatusCode int                    `json:"-"`
	Message        string                 `json:"message"`
	RequestID      string                 `json:"request-id"`
	Body           map[string]interface{} `json:"body,omitempty"`
}

func newBaseRequest(r *http.Request, message string, code int, kvs []interface{}) BaseResponse {
	return BaseResponse{
		HTTPStatusCode: code,
		Message:        message,
	}
}
