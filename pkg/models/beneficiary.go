package models

import (
	"time"
)

// Beneficiary represents a beneficiary
type Beneficiary struct {
	ReferredBy     string           `json:"referred_by,omitempty"`
	PartnerId      string           `json:"partner_id,omitempty"`
	AuditInfo      *AuditInfo       `json:"audit_info,omitempty"`
	BeneficiaryId  string           `json:"beneficiary_id"`
	UserId         string           `json:"user_id"`
	Type           *BeneficiaryType `json:"type"`
	Code           string           `json:"code"`
	Status         interface{}      `json:"status"`
	KycStatus      interface{}      `json:"kyc_status"`
	KycCompletedAt time.Time        `json:"kyc_completed_at,omitempty"`
	ReferralCode   string           `json:"referral_code,omitempty"`
	RiskScore      string           `json:"risk_score,omitempty"`
}
