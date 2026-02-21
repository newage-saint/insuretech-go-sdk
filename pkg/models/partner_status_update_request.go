package models

// PartnerStatusUpdateRequest represents a partner_status_update_request
type PartnerStatusUpdateRequest struct {
	Reason    string `json:"reason,omitempty"`
	PartnerId string `json:"partner_id"`
	Status    string `json:"status,omitempty"`
}
