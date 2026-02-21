package models

// WorkflowDefinition represents a workflow_definition
type WorkflowDefinition struct {
	WorkflowDefinitionId string        `json:"workflow_definition_id"`
	Name                 string        `json:"name"`
	Description          string        `json:"description,omitempty"`
	Type                 *WorkflowType `json:"type"`
	Steps                string        `json:"steps"`
	Version              int           `json:"version"`
	EntityType           string        `json:"entity_type"`
	Conditions           string        `json:"conditions,omitempty"`
	Status               interface{}   `json:"status"`
	AuditInfo            interface{}   `json:"audit_info"`
}
