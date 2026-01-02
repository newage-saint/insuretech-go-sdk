package models

// RoleAssignmentRequest represents a role_assignment_request
type RoleAssignmentRequest struct {
	ExpiresInDays string `json:"expires_in_days,omitempty"`
	UserId        string `json:"user_id"`
	RoleId        string `json:"role_id"`
	TenantId      string `json:"tenant_id"`
}
