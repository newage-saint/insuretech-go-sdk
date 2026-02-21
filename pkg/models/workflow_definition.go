package models

// WorkflowDefinition represents a workflow_definition
type WorkflowDefinition struct {
	Description          string        `json:"description,omitempty"`
	Conditions           string        `json:"conditions,omitempty"`
	Version              int           `json:"version"`
	AuditInfo            interface{}   `json:"audit_info"`
	Name                 string        `json:"name"`
	Type                 *WorkflowType `json:"type"`
	EntityType           string        `json:"entity_type"`
	Steps                string        `json:"steps"`
	Status               interface{}   `json:"status"`
	WorkflowDefinitionId string        `json:"workflow_definition_id"`
}
