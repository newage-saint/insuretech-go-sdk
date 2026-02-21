package models

import (
	"time"
)

// MFSWebhook represents a mfs_webhook
type MFSWebhook struct {
	Provider         string      `json:"provider"`
	EventType        string      `json:"event_type"`
	Headers          string      `json:"headers,omitempty"`
	Status           interface{} `json:"status"`
	SignatureValid   bool        `json:"signature_valid,omitempty"`
	MfsTransactionId string      `json:"mfs_transaction_id,omitempty"`
	ProcessedAt      time.Time   `json:"processed_at,omitempty"`
	AuditInfo        interface{} `json:"audit_info"`
	Id               string      `json:"id"`
	Payload          string      `json:"payload"`
	ErrorMessage     string      `json:"error_message,omitempty"`
}
