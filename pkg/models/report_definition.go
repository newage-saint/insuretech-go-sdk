package models

// ReportDefinition represents a report_definition
type ReportDefinition struct {
	Id           string          `json:"id"`
	Name         string          `json:"name"`
	Category     *ReportCategory `json:"category"`
	Description  string          `json:"description,omitempty"`
	Parameters   string          `json:"parameters,omitempty"`
	IsActive     bool            `json:"is_active,omitempty"`
	Query        string          `json:"query"`
	FormatConfig string          `json:"format_config,omitempty"`
	AuditInfo    interface{}     `json:"audit_info"`
}
