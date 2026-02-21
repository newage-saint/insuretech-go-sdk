package models

import (
	"time"
)

// Endorsement represents a endorsement
type Endorsement struct {
	ApprovedBy            string           `json:"approved_by,omitempty"`
	Id                    string           `json:"id"`
	EndorsementNumber     string           `json:"endorsement_number"`
	PolicyId              string           `json:"policy_id"`
	Status                interface{}      `json:"status"`
	EffectiveDate         time.Time        `json:"effective_date"`
	ApprovedAt            time.Time        `json:"approved_at,omitempty"`
	AuditInfo             interface{}      `json:"audit_info"`
	Type                  *EndorsementType `json:"type"`
	Reason                string           `json:"reason"`
	Changes               string           `json:"changes"`
	PremiumAdjustment     *Money           `json:"premium_adjustment,omitempty"`
	PremiumRefundRequired bool             `json:"premium_refund_required,omitempty"`
	RequestedBy           string           `json:"requested_by"`
}
