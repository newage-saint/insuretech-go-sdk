package models

import (
	"time"
)

// OTPSentEvent represents a otp_sent_event
type OTPSentEvent struct {
	Provider          string    `json:"provider,omitempty"`
	SenderId          string    `json:"sender_id,omitempty"`
	ProviderMessageId string    `json:"provider_message_id,omitempty"`
	EventId           string    `json:"event_id,omitempty"`
	OtpId             string    `json:"otp_id,omitempty"`
	Timestamp         time.Time `json:"timestamp,omitempty"`
	MaskingUsed       bool      `json:"masking_used,omitempty"`
	Recipient         string    `json:"recipient,omitempty"`
	Type              string    `json:"type,omitempty"`
	Channel           string    `json:"channel,omitempty"`
}
