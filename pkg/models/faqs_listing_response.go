package models

// FAQsListingResponse represents a faqs_listing_response
type FAQsListingResponse struct {
	TotalCount    int    `json:"total_count,omitempty"`
	Error         *Error `json:"error,omitempty"`
	Faqs          []*FAQ `json:"faqs,omitempty"`
	NextPageToken string `json:"next_page_token,omitempty"`
}
