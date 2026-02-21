package models

import (
	"time"
)

// Notification represents a notification
type Notification struct {
	NotificationId string                 `json:"notification_id"`
	Channel        *NotificationChannel   `json:"channel"`
	TemplateData   map[string]interface{} `json:"template_data,omitempty"`
	ScheduledAt    time.Time              `json:"scheduled_at,omitempty"`
	DeliveredAt    time.Time              `json:"delivered_at,omitempty"`
	RetryCount     int                    `json:"retry_count"`
	Priority       interface{}            `json:"priority"`
	Status         interface{}            `json:"status"`
	SentAt         time.Time              `json:"sent_at,omitempty"`
	ReadAt         time.Time              `json:"read_at,omitempty"`
	CreatedAt      time.Time              `json:"created_at"`
	RecipientId    string                 `json:"recipient_id"`
	Subject        string                 `json:"subject,omitempty"`
	Message        string                 `json:"message"`
	ErrorMessage   string                 `json:"error_message,omitempty"`
	Type           *NotificationType      `json:"type"`
}
