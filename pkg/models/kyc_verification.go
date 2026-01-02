package models

import (
	"time"
)

// KYCVerification represents a kyc_verification
type KYCVerification struct {
	Provider           string              `json:"provider,omitempty"`
	VerificationResult string              `json:"verification_result,omitempty"`
	VerifiedBy         string              `json:"verified_by,omitempty"`
	ExpiresAt          time.Time           `json:"expires_at,omitempty"`
	AuditInfo          *AuditInfo          `json:"audit_info,omitempty"`
	RejectionReason    string              `json:"rejection_reason,omitempty"`
	Type               *VerificationType   `json:"type"`
	EntityType         string              `json:"entity_type"`
	Method             *VerificationMethod `json:"method"`
	ProviderReference  string              `json:"provider_reference,omitempty"`
	Documents          string              `json:"documents,omitempty"`
	Status             interface{}         `json:"status"`
	VerifiedAt         time.Time           `json:"verified_at,omitempty"`
	Id                 string              `json:"id"`
	EntityId           string              `json:"entity_id"`
}
