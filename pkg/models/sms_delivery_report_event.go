package models

import (
	"time"
)

// SMSDeliveryReportEvent represents a sms_delivery_report_event
type SMSDeliveryReportEvent struct {
	EventId           string    `json:"event_id,omitempty"`
	OtpId             string    `json:"otp_id,omitempty"`
	DeliveredAt       time.Time `json:"delivered_at,omitempty"`
	Carrier           string    `json:"carrier,omitempty"`
	ProviderMessageId string    `json:"provider_message_id,omitempty"`
	Msisdn            string    `json:"msisdn,omitempty"`
	Status            string    `json:"status,omitempty"`
	ErrorCode         string    `json:"error_code,omitempty"`
	Timestamp         time.Time `json:"timestamp,omitempty"`
}
