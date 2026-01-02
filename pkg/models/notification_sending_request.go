package models

// NotificationSendingRequest represents a notification_sending_request
type NotificationSendingRequest struct {
	ScheduleAfterSeconds string                 `json:"schedule_after_seconds,omitempty"`
	Type                 *NotificationType      `json:"type"`
	Subject              string                 `json:"subject,omitempty"`
	TemplateData         map[string]interface{} `json:"template_data,omitempty"`
	RecipientId          string                 `json:"recipient_id"`
	Channel              *NotificationChannel   `json:"channel,omitempty"`
	Message              string                 `json:"message,omitempty"`
	TemplateId           string                 `json:"template_id"`
	Priority             *NotificationPriority  `json:"priority,omitempty"`
}
