package models

// DocumentTemplateDeactivationResponse represents a document_template_deactivation_response
type DocumentTemplateDeactivationResponse struct {
	Error   *Error `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}
