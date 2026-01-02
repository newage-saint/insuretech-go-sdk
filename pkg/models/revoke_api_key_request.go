package models

// RevokeApiKeyRequest represents a revoke_api_key_request
type RevokeApiKeyRequest struct {
	ApiKeyId string `json:"api_key_id"`
	Reason   string `json:"reason,omitempty"`
}
