package models

// WorkflowDefinitionCreationResponse represents a workflow_definition_creation_response
type WorkflowDefinitionCreationResponse struct {
	Error                *Error `json:"error,omitempty"`
	WorkflowDefinitionId string `json:"workflow_definition_id,omitempty"`
	Message              string `json:"message,omitempty"`
}
