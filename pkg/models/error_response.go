package models

// ErrorResponse represents a error_response
type ErrorResponse struct {
	Message string                 `json:"message,omitempty"`
	Details map[string]interface{} `json:"details,omitempty"`
	Code    *WebrtcErrorCode       `json:"code,omitempty"`
}
