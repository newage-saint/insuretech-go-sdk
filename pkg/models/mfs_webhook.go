package models

import (
	"time"
)

// MFSWebhook represents a mfs_webhook
type MFSWebhook struct {
	MfsTransactionId string      `json:"mfs_transaction_id,omitempty"`
	ErrorMessage     string      `json:"error_message,omitempty"`
	ProcessedAt      time.Time   `json:"processed_at,omitempty"`
	Provider         string      `json:"provider"`
	Headers          string      `json:"headers,omitempty"`
	SignatureValid   bool        `json:"signature_valid,omitempty"`
	AuditInfo        interface{} `json:"audit_info"`
	Id               string      `json:"id"`
	EventType        string      `json:"event_type"`
	Payload          string      `json:"payload"`
	Status           interface{} `json:"status"`
}
