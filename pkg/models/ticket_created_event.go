package models

import (
	"time"
)

// TicketCreatedEvent represents a ticket_created_event
type TicketCreatedEvent struct {
	Priority      string    `json:"priority,omitempty"`
	CorrelationId string    `json:"correlation_id,omitempty"`
	TicketId      string    `json:"ticket_id,omitempty"`
	Type          string    `json:"type,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
	EventId       string    `json:"event_id,omitempty"`
	TicketNumber  string    `json:"ticket_number,omitempty"`
	BeneficiaryId string    `json:"beneficiary_id,omitempty"`
	Category      string    `json:"category,omitempty"`
}
