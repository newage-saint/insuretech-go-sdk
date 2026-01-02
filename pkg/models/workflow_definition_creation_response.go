package models

// WorkflowDefinitionCreationResponse represents a workflow_definition_creation_response
type WorkflowDefinitionCreationResponse struct {
	Message              string `json:"message,omitempty"`
	Error                *Error `json:"error,omitempty"`
	WorkflowDefinitionId string `json:"workflow_definition_id,omitempty"`
}
