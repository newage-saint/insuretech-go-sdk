package models

// NotificationSendingRequest represents a notification_sending_request
type NotificationSendingRequest struct {
	Channel              *NotificationChannel   `json:"channel,omitempty"`
	TemplateId           string                 `json:"template_id"`
	Priority             *NotificationPriority  `json:"priority,omitempty"`
	RecipientId          string                 `json:"recipient_id"`
	Type                 *NotificationType      `json:"type"`
	Subject              string                 `json:"subject,omitempty"`
	Message              string                 `json:"message,omitempty"`
	TemplateData         map[string]interface{} `json:"template_data,omitempty"`
	ScheduleAfterSeconds string                 `json:"schedule_after_seconds,omitempty"`
}
