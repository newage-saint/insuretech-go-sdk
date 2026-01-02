package models

// TaskUpdateRequest represents a task_update_request
type TaskUpdateRequest struct {
	Status      string `json:"status,omitempty"`
	Priority    string `json:"priority,omitempty"`
	DueDate     string `json:"due_date,omitempty"`
	TaskId      string `json:"task_id"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}
