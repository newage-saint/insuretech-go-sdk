package models

import (
	"time"
)

// TicketAssignedEvent represents a ticket_assigned_event
type TicketAssignedEvent struct {
	CorrelationId string    `json:"correlation_id,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
	EventId       string    `json:"event_id,omitempty"`
	TicketId      string    `json:"ticket_id,omitempty"`
	TicketNumber  string    `json:"ticket_number,omitempty"`
	AssignedTo    string    `json:"assigned_to,omitempty"`
}
