package models

// DocumentTemplate represents a document_template
type DocumentTemplate struct {
	Version         int                   `json:"version"`
	IsActive        bool                  `json:"is_active,omitempty"`
	Description     string                `json:"description,omitempty"`
	AuditInfo       interface{}           `json:"audit_info"`
	Id              string                `json:"id"`
	Name            string                `json:"name"`
	Type            *DocumentDocumentType `json:"type"`
	TemplateContent string                `json:"template_content"`
	OutputFormat    *OutputFormat         `json:"output_format"`
	Variables       string                `json:"variables,omitempty"`
}
