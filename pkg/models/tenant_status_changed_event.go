package models

import (
	"time"
)

// TenantStatusChangedEvent represents a tenant_status_changed_event
type TenantStatusChangedEvent struct {
	TenantId      string    `json:"tenant_id,omitempty"`
	OldStatus     string    `json:"old_status,omitempty"`
	NewStatus     string    `json:"new_status,omitempty"`
	CorrelationId string    `json:"correlation_id,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
	EventId       string    `json:"event_id,omitempty"`
}
