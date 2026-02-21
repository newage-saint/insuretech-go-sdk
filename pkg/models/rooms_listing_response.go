package models

// RoomsListingResponse represents a rooms_listing_response
type RoomsListingResponse struct {
	NextPageToken string  `json:"next_page_token,omitempty"`
	TotalCount    int     `json:"total_count,omitempty"`
	Rooms         []*Room `json:"rooms,omitempty"`
}
