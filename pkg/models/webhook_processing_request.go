package models

// WebhookProcessingRequest represents a webhook_processing_request
type WebhookProcessingRequest struct {
	IdempotencyKey string                 `json:"idempotency_key,omitempty"`
	Provider       string                 `json:"provider"`
	Payload        map[string]interface{} `json:"payload,omitempty"`
	Headers        map[string]interface{} `json:"headers,omitempty"`
}
