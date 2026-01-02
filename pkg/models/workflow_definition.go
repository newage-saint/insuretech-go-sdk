package models

// WorkflowDefinition represents a workflow_definition
type WorkflowDefinition struct {
	Conditions  string        `json:"conditions,omitempty"`
	AuditInfo   *AuditInfo    `json:"audit_info,omitempty"`
	Id          string        `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description,omitempty"`
	EntityType  string        `json:"entity_type"`
	Steps       string        `json:"steps"`
	Version     int           `json:"version"`
	Status      interface{}   `json:"status"`
	Type        *WorkflowType `json:"type"`
}
