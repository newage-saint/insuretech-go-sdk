package models

// RevokeAllSessionsRequest represents a revoke_all_sessions_request
type RevokeAllSessionsRequest struct {
	Reason                string `json:"reason,omitempty"`
	ExcludeCurrentSession bool   `json:"exclude_current_session,omitempty"`
	UserId                string `json:"user_id"`
}
