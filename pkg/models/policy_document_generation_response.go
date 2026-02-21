package models

// PolicyDocumentGenerationResponse represents a policy_document_generation_response
type PolicyDocumentGenerationResponse struct {
	Error       *Error `json:"error,omitempty"`
	DocumentUrl string `json:"document_url,omitempty"`
	QrCode      string `json:"qr_code,omitempty"`
}
