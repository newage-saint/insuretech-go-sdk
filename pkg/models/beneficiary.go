package models

import (
	"time"
)

// Beneficiary represents a beneficiary
type Beneficiary struct {
	BeneficiaryId  string           `json:"beneficiary_id"`
	UserId         string           `json:"user_id"`
	Type           *BeneficiaryType `json:"type"`
	Status         interface{}      `json:"status"`
	KycCompletedAt time.Time        `json:"kyc_completed_at,omitempty"`
	ReferralCode   string           `json:"referral_code,omitempty"`
	PartnerId      string           `json:"partner_id,omitempty"`
	Code           string           `json:"code"`
	KycStatus      interface{}      `json:"kyc_status"`
	RiskScore      string           `json:"risk_score,omitempty"`
	ReferredBy     string           `json:"referred_by,omitempty"`
	AuditInfo      interface{}      `json:"audit_info"`
}
