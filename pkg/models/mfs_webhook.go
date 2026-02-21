package models

import (
	"time"
)

// MFSWebhook represents a mfs_webhook
type MFSWebhook struct {
	Provider         string      `json:"provider"`
	EventType        string      `json:"event_type"`
	Payload          string      `json:"payload"`
	Headers          string      `json:"headers,omitempty"`
	SignatureValid   bool        `json:"signature_valid,omitempty"`
	ErrorMessage     string      `json:"error_message,omitempty"`
	Status           interface{} `json:"status"`
	MfsTransactionId string      `json:"mfs_transaction_id,omitempty"`
	ProcessedAt      time.Time   `json:"processed_at,omitempty"`
	AuditInfo        interface{} `json:"audit_info"`
	Id               string      `json:"id"`
}
