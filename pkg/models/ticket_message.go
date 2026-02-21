package models

// TicketMessage represents a ticket_message
type TicketMessage struct {
	Id          string       `json:"id"`
	TicketId    string       `json:"ticket_id"`
	SenderId    string       `json:"sender_id"`
	Type        *MessageType `json:"type"`
	Content     string       `json:"content"`
	Attachments string       `json:"attachments,omitempty"`
	IsInternal  bool         `json:"is_internal,omitempty"`
	AuditInfo   interface{}  `json:"audit_info"`
}
