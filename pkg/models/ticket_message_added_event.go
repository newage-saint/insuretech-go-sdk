package models

import (
	"time"
)

// TicketMessageAddedEvent represents a ticket_message_added_event
type TicketMessageAddedEvent struct {
	MessageId     string    `json:"message_id,omitempty"`
	TicketId      string    `json:"ticket_id,omitempty"`
	SenderId      string    `json:"sender_id,omitempty"`
	IsInternal    bool      `json:"is_internal,omitempty"`
	CorrelationId string    `json:"correlation_id,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
	EventId       string    `json:"event_id,omitempty"`
}
