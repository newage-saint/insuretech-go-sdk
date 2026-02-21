package models

// ErrorInfo represents a error_info
type ErrorInfo struct {
	Meta    map[string]interface{} `json:"meta,omitempty"`
	Code    int                    `json:"code,omitempty"`
	Message string                 `json:"message,omitempty"`
}
