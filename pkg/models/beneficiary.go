package models

import (
	"time"
)

// Beneficiary represents a beneficiary
type Beneficiary struct {
	Type           *BeneficiaryType `json:"type"`
	Code           string           `json:"code"`
	KycStatus      interface{}      `json:"kyc_status"`
	KycCompletedAt time.Time        `json:"kyc_completed_at,omitempty"`
	RiskScore      string           `json:"risk_score,omitempty"`
	ReferralCode   string           `json:"referral_code,omitempty"`
	AuditInfo      interface{}      `json:"audit_info"`
	BeneficiaryId  string           `json:"beneficiary_id"`
	UserId         string           `json:"user_id"`
	Status         interface{}      `json:"status"`
	ReferredBy     string           `json:"referred_by,omitempty"`
	PartnerId      string           `json:"partner_id,omitempty"`
}
