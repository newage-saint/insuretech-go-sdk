package models

import (
	"time"
)

// NotificationTemplate represents a notification_template
type NotificationTemplate struct {
	TemplateName    string               `json:"template_name"`
	Type            *NotificationType    `json:"type"`
	Channel         *NotificationChannel `json:"channel"`
	SubjectTemplate string               `json:"subject_template,omitempty"`
	BodyTemplate    string               `json:"body_template"`
	TemplateId      string               `json:"template_id"`
	Language        string               `json:"language"`
	CreatedAt       time.Time            `json:"created_at"`
	UpdatedAt       time.Time            `json:"updated_at"`
	IsActive        bool                 `json:"is_active"`
}
