package models

import (
	"time"
)

// PartnerAPIKeyRotationResponse represents a partner_api_key_rotation_response
type PartnerAPIKeyRotationResponse struct {
	NewApiKey    string    `json:"new_api_key,omitempty"`
	NewApiSecret string    `json:"new_api_secret,omitempty"`
	ExpiresAt    time.Time `json:"expires_at,omitempty"`
	Error        *Error    `json:"error,omitempty"`
}
