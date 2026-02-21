package models

// ReportScheduleCreationRequest represents a report_schedule_creation_request
type ReportScheduleCreationRequest struct {
	Frequency          string                 `json:"frequency,omitempty"`
	CronExpression     string                 `json:"cron_expression,omitempty"`
	Parameters         map[string]interface{} `json:"parameters,omitempty"`
	Recipients         []string               `json:"recipients,omitempty"`
	ReportDefinitionId string                 `json:"report_definition_id"`
	Name               string                 `json:"name"`
}
