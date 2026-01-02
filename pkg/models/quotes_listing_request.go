package models

// QuotesListingRequest represents a quotes_listing_request
type QuotesListingRequest struct {
	Status        string `json:"status,omitempty"`
	Page          int    `json:"page,omitempty"`
	PageSize      int    `json:"page_size,omitempty"`
	BeneficiaryId string `json:"beneficiary_id"`
}
