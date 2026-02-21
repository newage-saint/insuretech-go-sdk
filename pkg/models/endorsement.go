package models

import (
	"time"
)

// Endorsement represents a endorsement
type Endorsement struct {
	PolicyId              string           `json:"policy_id"`
	Type                  *EndorsementType `json:"type"`
	Changes               string           `json:"changes"`
	EffectiveDate         time.Time        `json:"effective_date"`
	AuditInfo             interface{}      `json:"audit_info"`
	Id                    string           `json:"id"`
	Reason                string           `json:"reason"`
	PremiumAdjustment     *Money           `json:"premium_adjustment,omitempty"`
	PremiumRefundRequired bool             `json:"premium_refund_required,omitempty"`
	Status                interface{}      `json:"status"`
	RequestedBy           string           `json:"requested_by"`
	ApprovedBy            string           `json:"approved_by,omitempty"`
	ApprovedAt            time.Time        `json:"approved_at,omitempty"`
	EndorsementNumber     string           `json:"endorsement_number"`
}
