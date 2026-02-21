package models

import (
	"time"
)

// Beneficiary represents a beneficiary
type Beneficiary struct {
	Type           *BeneficiaryType `json:"type"`
	Code           string           `json:"code"`
	Status         interface{}      `json:"status"`
	KycStatus      interface{}      `json:"kyc_status"`
	KycCompletedAt time.Time        `json:"kyc_completed_at,omitempty"`
	RiskScore      string           `json:"risk_score,omitempty"`
	ReferralCode   string           `json:"referral_code,omitempty"`
	ReferredBy     string           `json:"referred_by,omitempty"`
	BeneficiaryId  string           `json:"beneficiary_id"`
	UserId         string           `json:"user_id"`
	PartnerId      string           `json:"partner_id,omitempty"`
	AuditInfo      interface{}      `json:"audit_info"`
}
