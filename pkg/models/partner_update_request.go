package models

// PartnerUpdateRequest represents a partner_update_request
type PartnerUpdateRequest struct {
	PartnerId  string   `json:"partner_id"`
	Partner    *Partner `json:"partner,omitempty"`
	UpdateMask []string `json:"update_mask,omitempty"`
}
