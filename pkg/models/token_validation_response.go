package models

// TokenValidationResponse represents a token_validation_response
type TokenValidationResponse struct {
	Valid       bool     `json:"valid,omitempty"`
	UserId      string   `json:"user_id,omitempty"`
	Roles       []string `json:"roles,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
	Error       *Error   `json:"error,omitempty"`
}
