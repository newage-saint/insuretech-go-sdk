package models

// SessionsListingResponse represents a sessions_listing_response
type SessionsListingResponse struct {
	NextPageToken string     `json:"next_page_token,omitempty"`
	TotalCount    int        `json:"total_count,omitempty"`
	Error         *Error     `json:"error,omitempty"`
	Sessions      []*Session `json:"sessions,omitempty"`
}
