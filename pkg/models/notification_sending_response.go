package models

// NotificationSendingResponse represents a notification_sending_response
type NotificationSendingResponse struct {
	Message        string `json:"message,omitempty"`
	Error          *Error `json:"error,omitempty"`
	NotificationId string `json:"notification_id,omitempty"`
}
