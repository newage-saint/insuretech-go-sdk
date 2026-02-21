package models

import (
	"time"
)

// RenewalSchedule represents a renewal_schedule
type RenewalSchedule struct {
	Status          interface{}  `json:"status"`
	GracePeriodDays int          `json:"grace_period_days,omitempty"`
	GracePeriodEnd  time.Time    `json:"grace_period_end,omitempty"`
	RenewedPolicyId string       `json:"renewed_policy_id,omitempty"`
	AuditInfo       interface{}  `json:"audit_info"`
	PolicyId        string       `json:"policy_id"`
	RenewalDueDate  time.Time    `json:"renewal_due_date"`
	RenewalPremium  *Money       `json:"renewal_premium,omitempty"`
	RenewedAt       time.Time    `json:"renewed_at,omitempty"`
	Id              string       `json:"id"`
	RenewalType     *RenewalType `json:"renewal_type"`
}
