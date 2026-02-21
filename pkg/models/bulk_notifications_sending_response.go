package models

// BulkNotificationsSendingResponse represents a bulk_notifications_sending_response
type BulkNotificationsSendingResponse struct {
	SuccessCount    int      `json:"success_count,omitempty"`
	FailedCount     int      `json:"failed_count,omitempty"`
	Error           *Error   `json:"error,omitempty"`
	NotificationIds []string `json:"notification_ids,omitempty"`
}
