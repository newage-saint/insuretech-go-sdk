package models

// WorkflowTaskCompletionResponse represents a workflow_task_completion_response
type WorkflowTaskCompletionResponse struct {
	Message string `json:"message,omitempty"`
	Error   *Error `json:"error,omitempty"`
}
