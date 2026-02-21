package models

// RoomsListingRequest represents a rooms_listing_request
type RoomsListingRequest struct {
	StateFilter *RoomState `json:"state_filter,omitempty"`
	PageSize    int        `json:"page_size"`
	PageToken   string     `json:"page_token,omitempty"`
}
