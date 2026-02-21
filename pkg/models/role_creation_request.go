package models

// RoleCreationRequest represents a role_creation_request
type RoleCreationRequest struct {
	Description string   `json:"description,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
	TenantId    string   `json:"tenant_id"`
	RoleName    string   `json:"role_name,omitempty"`
	DisplayName string   `json:"display_name,omitempty"`
}
