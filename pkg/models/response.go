package models

// Response represents a response
type Response struct {
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	Success   bool                   `json:"success,omitempty"`
	Message   string                 `json:"message,omitempty"`
	ErrorCode string                 `json:"error_code,omitempty"`
}
