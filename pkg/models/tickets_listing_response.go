package models

// TicketsListingResponse represents a tickets_listing_response
type TicketsListingResponse struct {
	TotalCount int       `json:"total_count,omitempty"`
	Error      *Error    `json:"error,omitempty"`
	Tickets    []*Ticket `json:"tickets,omitempty"`
}
