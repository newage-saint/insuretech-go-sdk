package models

// ApiKeyRotationResponse represents a api_key_rotation_response
type ApiKeyRotationResponse struct {
	Error     *Error `json:"error,omitempty"`
	NewApiKey string `json:"new_api_key,omitempty"`
	Message   string `json:"message,omitempty"`
}
