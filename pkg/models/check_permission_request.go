package models

// CheckPermissionRequest represents a check_permission_request
type CheckPermissionRequest struct {
	Action   string                 `json:"action"`
	Context  map[string]interface{} `json:"context,omitempty"`
	UserId   string                 `json:"user_id"`
	Resource string                 `json:"resource,omitempty"`
}
