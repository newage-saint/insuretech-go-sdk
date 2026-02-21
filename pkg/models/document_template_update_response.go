package models

// DocumentTemplateUpdateResponse represents a document_template_update_response
type DocumentTemplateUpdateResponse struct {
	Error   *Error `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}
