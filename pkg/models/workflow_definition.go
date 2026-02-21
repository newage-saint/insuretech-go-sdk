package models

// WorkflowDefinition represents a workflow_definition
type WorkflowDefinition struct {
	WorkflowDefinitionId string        `json:"workflow_definition_id"`
	Name                 string        `json:"name"`
	AuditInfo            interface{}   `json:"audit_info"`
	Description          string        `json:"description,omitempty"`
	Type                 *WorkflowType `json:"type"`
	EntityType           string        `json:"entity_type"`
	Steps                string        `json:"steps"`
	Conditions           string        `json:"conditions,omitempty"`
	Version              int           `json:"version"`
	Status               interface{}   `json:"status"`
}
