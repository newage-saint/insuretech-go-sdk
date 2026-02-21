package models

// DocumentTemplate represents a document_template
type DocumentTemplate struct {
	Id              string                `json:"id"`
	Description     string                `json:"description,omitempty"`
	TemplateContent string                `json:"template_content"`
	OutputFormat    *OutputFormat         `json:"output_format"`
	Variables       string                `json:"variables,omitempty"`
	AuditInfo       interface{}           `json:"audit_info"`
	Name            string                `json:"name"`
	Type            *DocumentDocumentType `json:"type"`
	Version         int                   `json:"version"`
	IsActive        bool                  `json:"is_active,omitempty"`
}
