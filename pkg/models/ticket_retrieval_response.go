package models

// TicketRetrievalResponse represents a ticket_retrieval_response
type TicketRetrievalResponse struct {
	Error    *Error           `json:"error,omitempty"`
	Ticket   *Ticket          `json:"ticket,omitempty"`
	Messages []*TicketMessage `json:"messages,omitempty"`
}
