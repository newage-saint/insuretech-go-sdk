package models

import (
	"time"
)

// Ticket represents a ticket
type Ticket struct {
	Id                string          `json:"id"`
	Status            interface{}     `json:"status"`
	AuditInfo         interface{}     `json:"audit_info"`
	TicketNumber      string          `json:"ticket_number"`
	Type              *TicketType     `json:"type"`
	Priority          interface{}     `json:"priority"`
	BeneficiaryId     string          `json:"beneficiary_id"`
	Subject           string          `json:"subject"`
	Description       string          `json:"description"`
	RelatedEntityType string          `json:"related_entity_type,omitempty"`
	AssignedTo        string          `json:"assigned_to,omitempty"`
	ClosedAt          time.Time       `json:"closed_at,omitempty"`
	Category          *TicketCategory `json:"category"`
	RelatedEntityId   string          `json:"related_entity_id,omitempty"`
	ResolvedAt        time.Time       `json:"resolved_at,omitempty"`
}
