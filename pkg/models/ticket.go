package models

import (
	"time"
)

// Ticket represents a ticket
type Ticket struct {
	Type              *TicketType     `json:"type"`
	Priority          interface{}     `json:"priority"`
	RelatedEntityId   string          `json:"related_entity_id,omitempty"`
	ResolvedAt        time.Time       `json:"resolved_at,omitempty"`
	Id                string          `json:"id"`
	TicketNumber      string          `json:"ticket_number"`
	Category          *TicketCategory `json:"category"`
	Description       string          `json:"description"`
	ClosedAt          time.Time       `json:"closed_at,omitempty"`
	AuditInfo         *AuditInfo      `json:"audit_info,omitempty"`
	AssignedTo        string          `json:"assigned_to,omitempty"`
	BeneficiaryId     string          `json:"beneficiary_id"`
	Subject           string          `json:"subject"`
	Status            interface{}     `json:"status"`
	RelatedEntityType string          `json:"related_entity_type,omitempty"`
}
