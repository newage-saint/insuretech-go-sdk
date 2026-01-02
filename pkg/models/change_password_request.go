package models

// ChangePasswordRequest represents a change_password_request
type ChangePasswordRequest struct {
	UserId      string `json:"user_id"`
	OldPassword string `json:"old_password,omitempty"`
	NewPassword string `json:"new_password,omitempty"`
}
