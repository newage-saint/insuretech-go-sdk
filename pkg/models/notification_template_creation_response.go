package models

// NotificationTemplateCreationResponse represents a notification_template_creation_response
type NotificationTemplateCreationResponse struct {
	Error      *Error `json:"error,omitempty"`
	TemplateId string `json:"template_id,omitempty"`
	Message    string `json:"message,omitempty"`
}
