package models

// NotificationTemplateDeactivationResponse represents a notification_template_deactivation_response
type NotificationTemplateDeactivationResponse struct {
	Error   *Error `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}
