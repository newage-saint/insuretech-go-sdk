package models

// PartnerUpdateResponse represents a partner_update_response
type PartnerUpdateResponse struct {
	Error   *Error   `json:"error,omitempty"`
	Partner *Partner `json:"partner,omitempty"`
}
