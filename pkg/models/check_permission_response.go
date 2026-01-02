package models

// CheckPermissionResponse represents a check_permission_response
type CheckPermissionResponse struct {
	Reason          string   `json:"reason,omitempty"`
	AppliedPolicies []string `json:"applied_policies,omitempty"`
	Error           *Error   `json:"error,omitempty"`
	Allowed         bool     `json:"allowed,omitempty"`
}
