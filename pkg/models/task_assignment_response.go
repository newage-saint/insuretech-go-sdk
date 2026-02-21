package models

// TaskAssignmentResponse represents a task_assignment_response
type TaskAssignmentResponse struct {
	Error   *Error `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}
