package models

import (
	"time"
)

// KYCVerification represents a kyc_verification
type KYCVerification struct {
	ProviderReference  string              `json:"provider_reference,omitempty"`
	Status             interface{}         `json:"status"`
	RejectionReason    string              `json:"rejection_reason,omitempty"`
	VerifiedBy         string              `json:"verified_by,omitempty"`
	Type               *VerificationType   `json:"type"`
	Documents          string              `json:"documents,omitempty"`
	VerificationResult string              `json:"verification_result,omitempty"`
	VerifiedAt         time.Time           `json:"verified_at,omitempty"`
	ExpiresAt          time.Time           `json:"expires_at,omitempty"`
	AuditInfo          interface{}         `json:"audit_info"`
	Id                 string              `json:"id"`
	EntityType         string              `json:"entity_type"`
	EntityId           string              `json:"entity_id"`
	Provider           string              `json:"provider,omitempty"`
	Method             *VerificationMethod `json:"method"`
}
