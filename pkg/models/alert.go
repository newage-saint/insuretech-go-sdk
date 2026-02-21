package models

import (
	"time"
)

// Alert represents a alert
type Alert struct {
	AlertId    string         `json:"alert_id,omitempty"`
	BusinessId string         `json:"business_id,omitempty"`
	AlertType  *AlertType     `json:"alert_type,omitempty"`
	Message    string         `json:"message,omitempty"`
	Count      int            `json:"count,omitempty"`
	IsRead     bool           `json:"is_read,omitempty"`
	ReadAt     time.Time      `json:"read_at,omitempty"`
	Title      string         `json:"title,omitempty"`
	Severity   *AlertSeverity `json:"severity,omitempty"`
	CreatedAt  time.Time      `json:"created_at,omitempty"`
}
