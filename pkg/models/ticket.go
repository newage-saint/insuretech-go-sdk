package models

import (
	"time"
)

// Ticket represents a ticket
type Ticket struct {
	RelatedEntityId   string          `json:"related_entity_id,omitempty"`
	TicketNumber      string          `json:"ticket_number"`
	Type              *TicketType     `json:"type"`
	Priority          interface{}     `json:"priority"`
	Description       string          `json:"description"`
	AssignedTo        string          `json:"assigned_to,omitempty"`
	AuditInfo         interface{}     `json:"audit_info"`
	BeneficiaryId     string          `json:"beneficiary_id"`
	Subject           string          `json:"subject"`
	RelatedEntityType string          `json:"related_entity_type,omitempty"`
	ResolvedAt        time.Time       `json:"resolved_at,omitempty"`
	Status            interface{}     `json:"status"`
	ClosedAt          time.Time       `json:"closed_at,omitempty"`
	Id                string          `json:"id"`
	Category          *TicketCategory `json:"category"`
}
