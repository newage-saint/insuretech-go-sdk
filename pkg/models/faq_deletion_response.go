package models

// FAQDeletionResponse represents a faq_deletion_response
type FAQDeletionResponse struct {
	Error   *Error `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}
