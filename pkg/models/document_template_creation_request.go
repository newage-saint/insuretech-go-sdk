package models

// DocumentTemplateCreationRequest represents a document_template_creation_request
type DocumentTemplateCreationRequest struct {
	Variables       []string `json:"variables,omitempty"`
	Name            string   `json:"name"`
	Type            string   `json:"type"`
	Description     string   `json:"description,omitempty"`
	TemplateContent string   `json:"template_content,omitempty"`
	OutputFormat    string   `json:"output_format,omitempty"`
}
