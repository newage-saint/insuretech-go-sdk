package models

// FAQUpdateResponse represents a faq_update_response
type FAQUpdateResponse struct {
	Error   *Error `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}
