package models

// RevokeApiKeyResponse represents a revoke_api_key_response
type RevokeApiKeyResponse struct {
	Message string `json:"message,omitempty"`
	Error   *Error `json:"error,omitempty"`
}
