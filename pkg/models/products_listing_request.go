package models

// ProductsListingRequest represents a products_listing_request
type ProductsListingRequest struct {
	PageSize int              `json:"page_size,omitempty"`
	Category *ProductCategory `json:"category"`
	Page     int              `json:"page,omitempty"`
}
