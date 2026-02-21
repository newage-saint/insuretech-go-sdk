package models

// Error represents a error
type Error struct {
	HttpStatusCode    int                    `json:"http_status_code,omitempty"`
	DocumentationUrl  string                 `json:"documentation_url,omitempty"`
	FieldViolations   []*FieldViolation      `json:"field_violations,omitempty"`
	Retryable         bool                   `json:"retryable,omitempty"`
	ErrorId           string                 `json:"error_id,omitempty"`
	Code              string                 `json:"code,omitempty"`
	Message           string                 `json:"message,omitempty"`
	Details           map[string]interface{} `json:"details,omitempty"`
	RetryAfterSeconds int                    `json:"retry_after_seconds,omitempty"`
}
