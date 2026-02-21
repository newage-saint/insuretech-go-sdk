package models

// LogoutRequest represents a logout_request
type LogoutRequest struct {
	LogoutReason string `json:"logout_reason,omitempty"`
	SessionId    string `json:"session_id"`
	AccessToken  string `json:"access_token,omitempty"`
}
