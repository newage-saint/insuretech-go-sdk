package models

import (
	"time"
)

// Endorsement represents a endorsement
type Endorsement struct {
	Reason                string           `json:"reason"`
	Status                interface{}      `json:"status"`
	ApprovedBy            string           `json:"approved_by,omitempty"`
	ApprovedAt            time.Time        `json:"approved_at,omitempty"`
	AuditInfo             interface{}      `json:"audit_info"`
	EndorsementNumber     string           `json:"endorsement_number"`
	Type                  *EndorsementType `json:"type"`
	Changes               string           `json:"changes"`
	PremiumAdjustment     *Money           `json:"premium_adjustment,omitempty"`
	PremiumRefundRequired bool             `json:"premium_refund_required,omitempty"`
	RequestedBy           string           `json:"requested_by"`
	EffectiveDate         time.Time        `json:"effective_date"`
	Id                    string           `json:"id"`
	PolicyId              string           `json:"policy_id"`
}
