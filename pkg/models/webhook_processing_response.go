package models

// WebhookProcessingResponse represents a webhook_processing_response
type WebhookProcessingResponse struct {
	Message string `json:"message,omitempty"`
	Success bool   `json:"success,omitempty"`
	Error   *Error `json:"error,omitempty"`
}
