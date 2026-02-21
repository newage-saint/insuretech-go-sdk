package models

import (
	"time"
)

// OTP represents a otp
type OTP struct {
	Attempts          int       `json:"attempts,omitempty"`
	Verified          bool      `json:"verified,omitempty"`
	DlrStatus         string    `json:"dlr_status,omitempty"`
	Carrier           string    `json:"carrier,omitempty"`
	Channel           string    `json:"channel,omitempty"`
	OtpHash           string    `json:"otp_hash,omitempty"`
	Purpose           string    `json:"purpose,omitempty"`
	DeviceType        string    `json:"device_type,omitempty"`
	VerifiedAt        time.Time `json:"verified_at,omitempty"`
	ProviderMessageId string    `json:"provider_message_id,omitempty"`
	Recipient         string    `json:"recipient,omitempty"`
	DlrReceivedAt     time.Time `json:"dlr_received_at,omitempty"`
	CreatedAt         time.Time `json:"created_at,omitempty"`
	OtpId             string    `json:"otp_id,omitempty"`
	UserId            string    `json:"user_id,omitempty"`
	IpAddress         string    `json:"ip_address,omitempty"`
	ExpiresAt         time.Time `json:"expires_at,omitempty"`
	SenderId          string    `json:"sender_id,omitempty"`
}
