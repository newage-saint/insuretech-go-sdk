package models

// AddTicketMessageRequest represents a add_ticket_message_request
type AddTicketMessageRequest struct {
	IsInternal  bool   `json:"is_internal,omitempty"`
	Attachments string `json:"attachments,omitempty"`
	TicketId    string `json:"ticket_id"`
	Content     string `json:"content,omitempty"`
}
