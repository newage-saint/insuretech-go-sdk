package models

import (
	"time"
)

// OTPSentEvent represents a otp_sent_event
type OTPSentEvent struct {
	OtpId     string    `json:"otp_id,omitempty"`
	Recipient string    `json:"recipient,omitempty"`
	Type      string    `json:"type,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
	Channel   string    `json:"channel,omitempty"`
	EventId   string    `json:"event_id,omitempty"`
}
