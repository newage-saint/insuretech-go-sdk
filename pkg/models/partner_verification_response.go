package models

import (
	"time"
)

// PartnerVerificationResponse represents a partner_verification_response
type PartnerVerificationResponse struct {
	VerifiedAt         time.Time `json:"verified_at,omitempty"`
	VerifiedBy         string    `json:"verified_by,omitempty"`
	Error              *Error    `json:"error,omitempty"`
	Verified           bool      `json:"verified,omitempty"`
	VerificationStatus string    `json:"verification_status,omitempty"`
}
