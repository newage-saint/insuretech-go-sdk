package models

import (
	"time"
)

// RenewalSchedule represents a renewal_schedule
type RenewalSchedule struct {
	RenewedPolicyId string       `json:"renewed_policy_id,omitempty"`
	Id              string       `json:"id"`
	RenewalPremium  *Money       `json:"renewal_premium,omitempty"`
	Status          interface{}  `json:"status"`
	GracePeriodEnd  time.Time    `json:"grace_period_end,omitempty"`
	RenewedAt       time.Time    `json:"renewed_at,omitempty"`
	AuditInfo       interface{}  `json:"audit_info"`
	PolicyId        string       `json:"policy_id"`
	RenewalDueDate  time.Time    `json:"renewal_due_date"`
	RenewalType     *RenewalType `json:"renewal_type"`
	GracePeriodDays int          `json:"grace_period_days,omitempty"`
}
