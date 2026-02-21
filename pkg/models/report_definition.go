package models

// ReportDefinition represents a report_definition
type ReportDefinition struct {
	IsActive     bool            `json:"is_active,omitempty"`
	Id           string          `json:"id"`
	Category     *ReportCategory `json:"category"`
	AuditInfo    interface{}     `json:"audit_info"`
	Name         string          `json:"name"`
	Description  string          `json:"description,omitempty"`
	Query        string          `json:"query"`
	Parameters   string          `json:"parameters,omitempty"`
	FormatConfig string          `json:"format_config,omitempty"`
}
