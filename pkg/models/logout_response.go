package models

// LogoutResponse represents a logout_response
type LogoutResponse struct {
	Message string `json:"message,omitempty"`
	Error   *Error `json:"error,omitempty"`
}
