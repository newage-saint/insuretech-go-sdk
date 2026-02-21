package models

import (
	"time"
)

// Endorsement represents a endorsement
type Endorsement struct {
	AuditInfo             interface{}      `json:"audit_info"`
	Id                    string           `json:"id"`
	PolicyId              string           `json:"policy_id"`
	Reason                string           `json:"reason"`
	Changes               string           `json:"changes"`
	PremiumAdjustment     *Money           `json:"premium_adjustment,omitempty"`
	Status                interface{}      `json:"status"`
	RequestedBy           string           `json:"requested_by"`
	ApprovedBy            string           `json:"approved_by,omitempty"`
	EndorsementNumber     string           `json:"endorsement_number"`
	Type                  *EndorsementType `json:"type"`
	PremiumRefundRequired bool             `json:"premium_refund_required,omitempty"`
	EffectiveDate         time.Time        `json:"effective_date"`
	ApprovedAt            time.Time        `json:"approved_at,omitempty"`
}
