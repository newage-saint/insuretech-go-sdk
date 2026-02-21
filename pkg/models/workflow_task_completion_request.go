package models

// WorkflowTaskCompletionRequest represents a workflow_task_completion_request
type WorkflowTaskCompletionRequest struct {
	Decision string `json:"decision,omitempty"`
	Comments string `json:"comments,omitempty"`
	TaskId   string `json:"task_id"`
}
