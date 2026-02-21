package models

// PeersListingResponse represents a peers_listing_response
type PeersListingResponse struct {
	TotalCount int     `json:"total_count,omitempty"`
	Peers      []*Peer `json:"peers,omitempty"`
}
