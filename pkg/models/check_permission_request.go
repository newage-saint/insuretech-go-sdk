package models

// CheckPermissionRequest represents a check_permission_request
type CheckPermissionRequest struct {
	UserId   string                 `json:"user_id"`
	Resource string                 `json:"resource,omitempty"`
	Action   string                 `json:"action"`
	Context  map[string]interface{} `json:"context,omitempty"`
}
