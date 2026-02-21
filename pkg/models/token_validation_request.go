package models

// TokenValidationRequest represents a token_validation_request
type TokenValidationRequest struct {
	SessionId   string `json:"session_id"`
	AccessToken string `json:"access_token,omitempty"`
}
