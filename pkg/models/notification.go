package models

import (
	"time"
)

// Notification represents a notification
type Notification struct {
	SentAt         time.Time              `json:"sent_at,omitempty"`
	DeliveredAt    time.Time              `json:"delivered_at,omitempty"`
	ReadAt         time.Time              `json:"read_at,omitempty"`
	Channel        *NotificationChannel   `json:"channel"`
	Message        string                 `json:"message"`
	TemplateData   map[string]interface{} `json:"template_data,omitempty"`
	RetryCount     int                    `json:"retry_count"`
	ErrorMessage   string                 `json:"error_message,omitempty"`
	CreatedAt      time.Time              `json:"created_at"`
	NotificationId string                 `json:"notification_id"`
	Subject        string                 `json:"subject,omitempty"`
	RecipientId    string                 `json:"recipient_id"`
	Status         interface{}            `json:"status"`
	Type           *NotificationType      `json:"type"`
	Priority       interface{}            `json:"priority"`
	ScheduledAt    time.Time              `json:"scheduled_at,omitempty"`
}
