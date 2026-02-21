package models

// PartnerCreationResponse represents a partner_creation_response
type PartnerCreationResponse struct {
	Partner   *Partner `json:"partner,omitempty"`
	Error     *Error   `json:"error,omitempty"`
	PartnerId string   `json:"partner_id,omitempty"`
}
