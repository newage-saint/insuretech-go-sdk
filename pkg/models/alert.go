package models

import (
	"time"
)

// Alert represents a alert
type Alert struct {
	AlertType  *AlertType     `json:"alert_type,omitempty"`
	Title      string         `json:"title,omitempty"`
	Message    string         `json:"message,omitempty"`
	Count      int            `json:"count,omitempty"`
	CreatedAt  time.Time      `json:"created_at,omitempty"`
	ReadAt     time.Time      `json:"read_at,omitempty"`
	Severity   *AlertSeverity `json:"severity,omitempty"`
	IsRead     bool           `json:"is_read,omitempty"`
	AlertId    string         `json:"alert_id,omitempty"`
	BusinessId string         `json:"business_id,omitempty"`
}
