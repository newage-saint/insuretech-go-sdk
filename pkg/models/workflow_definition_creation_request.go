package models

// WorkflowDefinitionCreationRequest represents a workflow_definition_creation_request
type WorkflowDefinitionCreationRequest struct {
	Type        string                 `json:"type"`
	EntityType  string                 `json:"entity_type"`
	Steps       map[string]interface{} `json:"steps,omitempty"`
	Name        string                 `json:"name"`
	Description string                 `json:"description,omitempty"`
}
