package models

// WebhookProcessingResponse represents a webhook_processing_response
type WebhookProcessingResponse struct {
	Error   *Error `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
	Success bool   `json:"success,omitempty"`
}
