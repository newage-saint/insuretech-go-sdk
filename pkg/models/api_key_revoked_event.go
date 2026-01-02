package models

import (
	"time"
)

// ApiKeyRevokedEvent represents a api_key_revoked_event
type ApiKeyRevokedEvent struct {
	Timestamp     time.Time `json:"timestamp,omitempty"`
	EventId       string    `json:"event_id,omitempty"`
	ApiKeyId      string    `json:"api_key_id,omitempty"`
	Reason        string    `json:"reason,omitempty"`
	CorrelationId string    `json:"correlation_id,omitempty"`
}
