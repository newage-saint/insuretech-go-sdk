package models

// InsurerProductsListingResponse represents a insurer_products_listing_response
type InsurerProductsListingResponse struct {
	InsurerProducts []*InsurerProduct `json:"insurer_products,omitempty"`
	TotalCount      int               `json:"total_count,omitempty"`
	Error           *Error            `json:"error,omitempty"`
}
