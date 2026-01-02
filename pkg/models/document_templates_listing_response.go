package models

// DocumentTemplatesListingResponse represents a document_templates_listing_response
type DocumentTemplatesListingResponse struct {
	NextPageToken string              `json:"next_page_token,omitempty"`
	TotalCount    int                 `json:"total_count,omitempty"`
	Error         *Error              `json:"error,omitempty"`
	Templates     []*DocumentTemplate `json:"templates,omitempty"`
}
