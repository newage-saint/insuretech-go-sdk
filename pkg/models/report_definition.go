package models

// ReportDefinition represents a report_definition
type ReportDefinition struct {
	IsActive     bool            `json:"is_active,omitempty"`
	AuditInfo    interface{}     `json:"audit_info"`
	Category     *ReportCategory `json:"category"`
	Query        string          `json:"query"`
	Id           string          `json:"id"`
	Name         string          `json:"name"`
	Description  string          `json:"description,omitempty"`
	Parameters   string          `json:"parameters,omitempty"`
	FormatConfig string          `json:"format_config,omitempty"`
}
