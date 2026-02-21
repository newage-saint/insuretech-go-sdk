package models

// ReportDefinition represents a report_definition
type ReportDefinition struct {
	Name         string          `json:"name"`
	Category     *ReportCategory `json:"category"`
	Parameters   string          `json:"parameters,omitempty"`
	FormatConfig string          `json:"format_config,omitempty"`
	Id           string          `json:"id"`
	Description  string          `json:"description,omitempty"`
	Query        string          `json:"query"`
	IsActive     bool            `json:"is_active,omitempty"`
	AuditInfo    interface{}     `json:"audit_info"`
}
