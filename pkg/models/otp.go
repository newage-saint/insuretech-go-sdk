package models

import (
	"time"
)

// OTP represents a otp
type OTP struct {
	SenderId          string    `json:"sender_id,omitempty"`
	OtpHash           string    `json:"otp_hash,omitempty"`
	Purpose           string    `json:"purpose,omitempty"`
	DeviceType        string    `json:"device_type,omitempty"`
	Verified          bool      `json:"verified,omitempty"`
	ProviderMessageId string    `json:"provider_message_id,omitempty"`
	DlrStatus         string    `json:"dlr_status,omitempty"`
	Carrier           string    `json:"carrier,omitempty"`
	OtpId             string    `json:"otp_id,omitempty"`
	UserId            string    `json:"user_id,omitempty"`
	CreatedAt         time.Time `json:"created_at,omitempty"`
	Channel           string    `json:"channel,omitempty"`
	DlrReceivedAt     time.Time `json:"dlr_received_at,omitempty"`
	Attempts          int       `json:"attempts,omitempty"`
	VerifiedAt        time.Time `json:"verified_at,omitempty"`
	Recipient         string    `json:"recipient,omitempty"`
	IpAddress         string    `json:"ip_address,omitempty"`
	ExpiresAt         time.Time `json:"expires_at,omitempty"`
}
