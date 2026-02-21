package models

// BeneficiariesListingRequest represents a beneficiaries_listing_request
type BeneficiariesListingRequest struct {
	Page     int    `json:"page,omitempty"`
	PageSize int    `json:"page_size,omitempty"`
	Type     string `json:"type"`
	Status   string `json:"status,omitempty"`
}
