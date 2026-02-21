package models

import (
	"time"
)

// SMSDeliveryReportEvent represents a sms_delivery_report_event
type SMSDeliveryReportEvent struct {
	ProviderMessageId string    `json:"provider_message_id,omitempty"`
	Status            string    `json:"status,omitempty"`
	ErrorCode         string    `json:"error_code,omitempty"`
	DeliveredAt       time.Time `json:"delivered_at,omitempty"`
	Timestamp         time.Time `json:"timestamp,omitempty"`
	EventId           string    `json:"event_id,omitempty"`
	OtpId             string    `json:"otp_id,omitempty"`
	Msisdn            string    `json:"msisdn,omitempty"`
	Carrier           string    `json:"carrier,omitempty"`
}
