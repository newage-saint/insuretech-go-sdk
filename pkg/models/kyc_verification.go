package models

import (
	"time"
)

// KYCVerification represents a kyc_verification
type KYCVerification struct {
	Type               *VerificationType   `json:"type"`
	EntityId           string              `json:"entity_id"`
	Documents          string              `json:"documents,omitempty"`
	RejectionReason    string              `json:"rejection_reason,omitempty"`
	Method             *VerificationMethod `json:"method"`
	ProviderReference  string              `json:"provider_reference,omitempty"`
	VerificationResult string              `json:"verification_result,omitempty"`
	VerifiedBy         string              `json:"verified_by,omitempty"`
	AuditInfo          interface{}         `json:"audit_info"`
	EntityType         string              `json:"entity_type"`
	Provider           string              `json:"provider,omitempty"`
	Status             interface{}         `json:"status"`
	VerifiedAt         time.Time           `json:"verified_at,omitempty"`
	Id                 string              `json:"id"`
	ExpiresAt          time.Time           `json:"expires_at,omitempty"`
}
