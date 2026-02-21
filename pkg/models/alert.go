package models

import (
	"time"
)

// Alert represents a alert
type Alert struct {
	AlertId    string         `json:"alert_id,omitempty"`
	BusinessId string         `json:"business_id,omitempty"`
	Title      string         `json:"title,omitempty"`
	Message    string         `json:"message,omitempty"`
	IsRead     bool           `json:"is_read,omitempty"`
	CreatedAt  time.Time      `json:"created_at,omitempty"`
	AlertType  *AlertType     `json:"alert_type,omitempty"`
	Severity   *AlertSeverity `json:"severity,omitempty"`
	Count      int            `json:"count,omitempty"`
	ReadAt     time.Time      `json:"read_at,omitempty"`
}
