package models

import (
	"time"
)

// KYCVerification represents a kyc_verification
type KYCVerification struct {
	Type               *VerificationType   `json:"type"`
	ExpiresAt          time.Time           `json:"expires_at,omitempty"`
	AuditInfo          interface{}         `json:"audit_info"`
	EntityType         string              `json:"entity_type"`
	EntityId           string              `json:"entity_id"`
	Method             *VerificationMethod `json:"method"`
	Status             interface{}         `json:"status"`
	RejectionReason    string              `json:"rejection_reason,omitempty"`
	VerifiedAt         time.Time           `json:"verified_at,omitempty"`
	ProviderReference  string              `json:"provider_reference,omitempty"`
	Documents          string              `json:"documents,omitempty"`
	VerificationResult string              `json:"verification_result,omitempty"`
	Id                 string              `json:"id"`
	Provider           string              `json:"provider,omitempty"`
	VerifiedBy         string              `json:"verified_by,omitempty"`
}
