package models

// RoleAssignmentResponse represents a role_assignment_response
type RoleAssignmentResponse struct {
	AssignmentId string `json:"assignment_id,omitempty"`
	Message      string `json:"message,omitempty"`
	Error        *Error `json:"error,omitempty"`
}
