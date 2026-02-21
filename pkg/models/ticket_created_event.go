package models

import (
	"time"
)

// TicketCreatedEvent represents a ticket_created_event
type TicketCreatedEvent struct {
	TicketNumber  string    `json:"ticket_number,omitempty"`
	BeneficiaryId string    `json:"beneficiary_id,omitempty"`
	Type          string    `json:"type,omitempty"`
	CorrelationId string    `json:"correlation_id,omitempty"`
	EventId       string    `json:"event_id,omitempty"`
	TicketId      string    `json:"ticket_id,omitempty"`
	Category      string    `json:"category,omitempty"`
	Priority      string    `json:"priority,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}
