package models

// Error represents a error
type Error struct {
	Code              string                 `json:"code,omitempty"`
	Message           string                 `json:"message,omitempty"`
	Details           map[string]interface{} `json:"details,omitempty"`
	FieldViolations   []*FieldViolation      `json:"field_violations,omitempty"`
	RetryAfterSeconds int                    `json:"retry_after_seconds,omitempty"`
	ErrorId           string                 `json:"error_id,omitempty"`
	Retryable         bool                   `json:"retryable,omitempty"`
	HttpStatusCode    int                    `json:"http_status_code,omitempty"`
	DocumentationUrl  string                 `json:"documentation_url,omitempty"`
}
