package models

import (
	"time"
)

// Endorsement represents a endorsement
type Endorsement struct {
	RequestedBy           string           `json:"requested_by"`
	ApprovedBy            string           `json:"approved_by,omitempty"`
	Type                  *EndorsementType `json:"type"`
	Reason                string           `json:"reason"`
	Status                interface{}      `json:"status"`
	EffectiveDate         time.Time        `json:"effective_date"`
	ApprovedAt            time.Time        `json:"approved_at,omitempty"`
	AuditInfo             *AuditInfo       `json:"audit_info,omitempty"`
	Id                    string           `json:"id"`
	EndorsementNumber     string           `json:"endorsement_number"`
	PolicyId              string           `json:"policy_id"`
	Changes               string           `json:"changes"`
	PremiumAdjustment     *Money           `json:"premium_adjustment,omitempty"`
	PremiumRefundRequired bool             `json:"premium_refund_required,omitempty"`
}
