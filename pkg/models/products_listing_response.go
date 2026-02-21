package models

// ProductsListingResponse represents a products_listing_response
type ProductsListingResponse struct {
	TotalCount int        `json:"total_count,omitempty"`
	Page       int        `json:"page,omitempty"`
	PageSize   int        `json:"page_size,omitempty"`
	Error      *Error     `json:"error,omitempty"`
	Products   []*Product `json:"products,omitempty"`
}
