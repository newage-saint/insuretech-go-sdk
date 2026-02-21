package models

import (
	"time"
)

// NotificationTemplate represents a notification_template
type NotificationTemplate struct {
	TemplateId      string               `json:"template_id"`
	Type            *NotificationType    `json:"type"`
	SubjectTemplate string               `json:"subject_template,omitempty"`
	UpdatedAt       time.Time            `json:"updated_at"`
	IsActive        bool                 `json:"is_active"`
	TemplateName    string               `json:"template_name"`
	Channel         *NotificationChannel `json:"channel"`
	BodyTemplate    string               `json:"body_template"`
	Language        string               `json:"language"`
	CreatedAt       time.Time            `json:"created_at"`
}
