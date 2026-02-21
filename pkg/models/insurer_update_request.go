package models

// InsurerUpdateRequest represents a insurer_update_request
type InsurerUpdateRequest struct {
	Email     string `json:"email"`
	Phone     string `json:"phone,omitempty"`
	Status    string `json:"status,omitempty"`
	InsurerId string `json:"insurer_id"`
	Name      string `json:"name"`
}
