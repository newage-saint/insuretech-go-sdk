package models

import (
	"time"
)

// Ticket represents a ticket
type Ticket struct {
	Id                string          `json:"id"`
	Type              *TicketType     `json:"type"`
	Priority          interface{}     `json:"priority"`
	RelatedEntityType string          `json:"related_entity_type,omitempty"`
	RelatedEntityId   string          `json:"related_entity_id,omitempty"`
	TicketNumber      string          `json:"ticket_number"`
	BeneficiaryId     string          `json:"beneficiary_id"`
	Category          *TicketCategory `json:"category"`
	Subject           string          `json:"subject"`
	Description       string          `json:"description"`
	Status            interface{}     `json:"status"`
	AssignedTo        string          `json:"assigned_to,omitempty"`
	AuditInfo         interface{}     `json:"audit_info"`
	ResolvedAt        time.Time       `json:"resolved_at,omitempty"`
	ClosedAt          time.Time       `json:"closed_at,omitempty"`
}
