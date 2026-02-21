package models

// ApiKeyGenerationResponse represents a api_key_generation_response
type ApiKeyGenerationResponse struct {
	ApiKey   string `json:"api_key,omitempty"`
	Message  string `json:"message,omitempty"`
	Error    *Error `json:"error,omitempty"`
	ApiKeyId string `json:"api_key_id,omitempty"`
}
