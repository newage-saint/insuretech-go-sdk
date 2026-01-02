package models

import (
	"time"
)

// OTP represents a otp
type OTP struct {
	Recipient  string    `json:"recipient"`
	Type       *OTPType  `json:"type"`
	Code       string    `json:"code"`
	ExpiresAt  time.Time `json:"expires_at"`
	Attempts   int       `json:"attempts"`
	Verified   bool      `json:"verified"`
	VerifiedAt time.Time `json:"verified_at,omitempty"`
	OtpId      string    `json:"otp_id"`
	CreatedAt  time.Time `json:"created_at"`
}
