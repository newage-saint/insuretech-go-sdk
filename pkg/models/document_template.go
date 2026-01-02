package models

// DocumentTemplate represents a document_template
type DocumentTemplate struct {
	Id              string                `json:"id"`
	Name            string                `json:"name"`
	OutputFormat    *OutputFormat         `json:"output_format"`
	Variables       string                `json:"variables,omitempty"`
	Version         int                   `json:"version"`
	IsActive        bool                  `json:"is_active,omitempty"`
	Type            *DocumentDocumentType `json:"type"`
	Description     string                `json:"description,omitempty"`
	TemplateContent string                `json:"template_content"`
	AuditInfo       *AuditInfo            `json:"audit_info,omitempty"`
}
