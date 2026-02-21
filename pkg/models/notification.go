package models

import (
	"time"
)

// Notification represents a notification
type Notification struct {
	Type           *NotificationType      `json:"type"`
	Channel        *NotificationChannel   `json:"channel"`
	ScheduledAt    time.Time              `json:"scheduled_at,omitempty"`
	RetryCount     int                    `json:"retry_count"`
	ErrorMessage   string                 `json:"error_message,omitempty"`
	Priority       interface{}            `json:"priority"`
	SentAt         time.Time              `json:"sent_at,omitempty"`
	DeliveredAt    time.Time              `json:"delivered_at,omitempty"`
	NotificationId string                 `json:"notification_id"`
	Subject        string                 `json:"subject,omitempty"`
	Status         interface{}            `json:"status"`
	CreatedAt      time.Time              `json:"created_at"`
	RecipientId    string                 `json:"recipient_id"`
	Message        string                 `json:"message"`
	TemplateData   map[string]interface{} `json:"template_data,omitempty"`
	ReadAt         time.Time              `json:"read_at,omitempty"`
}
