package models

// RoleCreationResponse represents a role_creation_response
type RoleCreationResponse struct {
	RoleId  string `json:"role_id,omitempty"`
	Message string `json:"message,omitempty"`
	Error   *Error `json:"error,omitempty"`
}
