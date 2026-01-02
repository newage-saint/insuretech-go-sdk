package models

// ApiKeyRotationResponse represents a api_key_rotation_response
type ApiKeyRotationResponse struct {
	NewApiKey string `json:"new_api_key,omitempty"`
	Message   string `json:"message,omitempty"`
	Error     *Error `json:"error,omitempty"`
}
