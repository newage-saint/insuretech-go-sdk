package models

// InsurerConfigUpdateResponse represents a insurer_config_update_response
type InsurerConfigUpdateResponse struct {
	Error   *Error `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}
