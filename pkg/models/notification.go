package models

import (
	"time"
)

// Notification represents a notification
type Notification struct {
	RetryCount     int                    `json:"retry_count"`
	NotificationId string                 `json:"notification_id"`
	RecipientId    string                 `json:"recipient_id"`
	Channel        *NotificationChannel   `json:"channel"`
	TemplateData   map[string]interface{} `json:"template_data,omitempty"`
	DeliveredAt    time.Time              `json:"delivered_at,omitempty"`
	ReadAt         time.Time              `json:"read_at,omitempty"`
	Message        string                 `json:"message"`
	SentAt         time.Time              `json:"sent_at,omitempty"`
	ErrorMessage   string                 `json:"error_message,omitempty"`
	Priority       interface{}            `json:"priority"`
	Status         interface{}            `json:"status"`
	Type           *NotificationType      `json:"type"`
	Subject        string                 `json:"subject,omitempty"`
	ScheduledAt    time.Time              `json:"scheduled_at,omitempty"`
	CreatedAt      time.Time              `json:"created_at"`
}
