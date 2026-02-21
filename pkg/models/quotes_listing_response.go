package models

// QuotesListingResponse represents a quotes_listing_response
type QuotesListingResponse struct {
	TotalCount int      `json:"total_count,omitempty"`
	Error      *Error   `json:"error,omitempty"`
	Quotes     []*Quote `json:"quotes,omitempty"`
}
