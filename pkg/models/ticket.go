package models

import (
	"time"
)

// Ticket represents a ticket
type Ticket struct {
	Id                string          `json:"id"`
	Priority          interface{}     `json:"priority"`
	Subject           string          `json:"subject"`
	ResolvedAt        time.Time       `json:"resolved_at,omitempty"`
	TicketNumber      string          `json:"ticket_number"`
	Category          *TicketCategory `json:"category"`
	Description       string          `json:"description"`
	RelatedEntityType string          `json:"related_entity_type,omitempty"`
	RelatedEntityId   string          `json:"related_entity_id,omitempty"`
	ClosedAt          time.Time       `json:"closed_at,omitempty"`
	AuditInfo         interface{}     `json:"audit_info"`
	BeneficiaryId     string          `json:"beneficiary_id"`
	Type              *TicketType     `json:"type"`
	Status            interface{}     `json:"status"`
	AssignedTo        string          `json:"assigned_to,omitempty"`
}
