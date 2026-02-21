package models

import (
	"time"
)

// Alert represents a alert
type Alert struct {
	Severity   *AlertSeverity `json:"severity,omitempty"`
	Count      int            `json:"count,omitempty"`
	IsRead     bool           `json:"is_read,omitempty"`
	CreatedAt  time.Time      `json:"created_at,omitempty"`
	ReadAt     time.Time      `json:"read_at,omitempty"`
	AlertId    string         `json:"alert_id,omitempty"`
	Message    string         `json:"message,omitempty"`
	BusinessId string         `json:"business_id,omitempty"`
	AlertType  *AlertType     `json:"alert_type,omitempty"`
	Title      string         `json:"title,omitempty"`
}
