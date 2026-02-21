package models

// InsurerProductsListingRequest represents a insurer_products_listing_request
type InsurerProductsListingRequest struct {
	PageSize  int    `json:"page_size,omitempty"`
	InsurerId string `json:"insurer_id"`
	Status    string `json:"status,omitempty"`
	Page      int    `json:"page,omitempty"`
}
