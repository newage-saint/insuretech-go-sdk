package models

import (
	"time"
)

// KYCVerification represents a kyc_verification
type KYCVerification struct {
	Type               *VerificationType   `json:"type"`
	ProviderReference  string              `json:"provider_reference,omitempty"`
	VerificationResult string              `json:"verification_result,omitempty"`
	ExpiresAt          time.Time           `json:"expires_at,omitempty"`
	Method             *VerificationMethod `json:"method"`
	RejectionReason    string              `json:"rejection_reason,omitempty"`
	Provider           string              `json:"provider,omitempty"`
	Documents          string              `json:"documents,omitempty"`
	Status             interface{}         `json:"status"`
	VerifiedBy         string              `json:"verified_by,omitempty"`
	Id                 string              `json:"id"`
	EntityType         string              `json:"entity_type"`
	EntityId           string              `json:"entity_id"`
	VerifiedAt         time.Time           `json:"verified_at,omitempty"`
	AuditInfo          interface{}         `json:"audit_info"`
}
