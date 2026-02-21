package models

// WorkflowDefinition represents a workflow_definition
type WorkflowDefinition struct {
	Steps                string        `json:"steps"`
	Conditions           string        `json:"conditions,omitempty"`
	Status               interface{}   `json:"status"`
	AuditInfo            interface{}   `json:"audit_info"`
	Type                 *WorkflowType `json:"type"`
	Version              int           `json:"version"`
	WorkflowDefinitionId string        `json:"workflow_definition_id"`
	Name                 string        `json:"name"`
	Description          string        `json:"description,omitempty"`
	EntityType           string        `json:"entity_type"`
}
