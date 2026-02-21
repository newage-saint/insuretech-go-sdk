package models

// ResetPasswordResponse represents a reset_password_response
type ResetPasswordResponse struct {
	Error   *Error `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}
