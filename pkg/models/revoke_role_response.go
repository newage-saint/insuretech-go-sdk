package models

// RevokeRoleResponse represents a revoke_role_response
type RevokeRoleResponse struct {
	Error   *Error `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}
