package models

import (
	"time"
)

// Notification represents a notification
type Notification struct {
	NotificationId string                 `json:"notification_id"`
	Subject        string                 `json:"subject,omitempty"`
	Message        string                 `json:"message"`
	Priority       interface{}            `json:"priority"`
	SentAt         time.Time              `json:"sent_at,omitempty"`
	CreatedAt      time.Time              `json:"created_at"`
	RecipientId    string                 `json:"recipient_id"`
	Channel        *NotificationChannel   `json:"channel"`
	TemplateData   map[string]interface{} `json:"template_data,omitempty"`
	ReadAt         time.Time              `json:"read_at,omitempty"`
	Status         interface{}            `json:"status"`
	ScheduledAt    time.Time              `json:"scheduled_at,omitempty"`
	DeliveredAt    time.Time              `json:"delivered_at,omitempty"`
	RetryCount     int                    `json:"retry_count"`
	ErrorMessage   string                 `json:"error_message,omitempty"`
	Type           *NotificationType      `json:"type"`
}
