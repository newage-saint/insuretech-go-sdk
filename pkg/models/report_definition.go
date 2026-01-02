package models

// ReportDefinition represents a report_definition
type ReportDefinition struct {
	Description  string          `json:"description,omitempty"`
	FormatConfig string          `json:"format_config,omitempty"`
	IsActive     bool            `json:"is_active,omitempty"`
	Id           string          `json:"id"`
	Name         string          `json:"name"`
	Query        string          `json:"query"`
	Parameters   string          `json:"parameters,omitempty"`
	AuditInfo    *AuditInfo      `json:"audit_info,omitempty"`
	Category     *ReportCategory `json:"category"`
}
