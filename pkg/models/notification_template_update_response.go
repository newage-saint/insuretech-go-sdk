package models

// NotificationTemplateUpdateResponse represents a notification_template_update_response
type NotificationTemplateUpdateResponse struct {
	Message string `json:"message,omitempty"`
	Error   *Error `json:"error,omitempty"`
}
