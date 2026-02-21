package models

import (
	"time"
)

// OTP represents a otp
type OTP struct {
	UserId            string    `json:"user_id,omitempty"`
	Purpose           string    `json:"purpose,omitempty"`
	DeviceType        string    `json:"device_type,omitempty"`
	CreatedAt         time.Time `json:"created_at,omitempty"`
	Channel           string    `json:"channel,omitempty"`
	Recipient         string    `json:"recipient,omitempty"`
	OtpId             string    `json:"otp_id,omitempty"`
	ProviderMessageId string    `json:"provider_message_id,omitempty"`
	SenderId          string    `json:"sender_id,omitempty"`
	Carrier           string    `json:"carrier,omitempty"`
	DlrReceivedAt     time.Time `json:"dlr_received_at,omitempty"`
	OtpHash           string    `json:"otp_hash,omitempty"`
	IpAddress         string    `json:"ip_address,omitempty"`
	ExpiresAt         time.Time `json:"expires_at,omitempty"`
	Attempts          int       `json:"attempts,omitempty"`
	VerifiedAt        time.Time `json:"verified_at,omitempty"`
	DlrStatus         string    `json:"dlr_status,omitempty"`
	Verified          bool      `json:"verified,omitempty"`
}
