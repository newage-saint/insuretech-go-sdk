package models

// BeneficiariesListingResponse represents a beneficiaries_listing_response
type BeneficiariesListingResponse struct {
	Beneficiaries []*Beneficiary `json:"beneficiaries,omitempty"`
	TotalCount    int            `json:"total_count,omitempty"`
	Error         *Error         `json:"error,omitempty"`
}
