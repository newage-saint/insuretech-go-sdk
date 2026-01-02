package models

// NotificationTemplateDeactivationRequest represents a notification_template_deactivation_request
type NotificationTemplateDeactivationRequest struct {
	TemplateId string `json:"template_id"`
	Reason     string `json:"reason,omitempty"`
}
