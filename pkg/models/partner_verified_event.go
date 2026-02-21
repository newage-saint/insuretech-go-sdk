package models

import (
	"time"
)

// PartnerVerifiedEvent represents a partner_verified_event
type PartnerVerifiedEvent struct {
	VerifiedBy string    `json:"verified_by,omitempty"`
	Timestamp  time.Time `json:"timestamp,omitempty"`
	EventId    string    `json:"event_id,omitempty"`
	PartnerId  string    `json:"partner_id,omitempty"`
}
