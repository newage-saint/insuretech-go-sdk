package models

import (
	"time"
)

// MFSWebhook represents a mfs_webhook
type MFSWebhook struct {
	EventType        string      `json:"event_type"`
	Headers          string      `json:"headers,omitempty"`
	SignatureValid   bool        `json:"signature_valid,omitempty"`
	ErrorMessage     string      `json:"error_message,omitempty"`
	ProcessedAt      time.Time   `json:"processed_at,omitempty"`
	AuditInfo        *AuditInfo  `json:"audit_info,omitempty"`
	Id               string      `json:"id"`
	Provider         string      `json:"provider"`
	Payload          string      `json:"payload"`
	Status           interface{} `json:"status"`
	MfsTransactionId string      `json:"mfs_transaction_id,omitempty"`
}
