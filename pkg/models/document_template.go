package models

// DocumentTemplate represents a document_template
type DocumentTemplate struct {
	Type            *DocumentDocumentType `json:"type"`
	Description     string                `json:"description,omitempty"`
	TemplateContent string                `json:"template_content"`
	OutputFormat    *OutputFormat         `json:"output_format"`
	Variables       string                `json:"variables,omitempty"`
	IsActive        bool                  `json:"is_active,omitempty"`
	AuditInfo       interface{}           `json:"audit_info"`
	Id              string                `json:"id"`
	Name            string                `json:"name"`
	Version         int                   `json:"version"`
}
