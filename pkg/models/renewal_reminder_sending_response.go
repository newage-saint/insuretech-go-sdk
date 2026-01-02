package models

// RenewalReminderSendingResponse represents a renewal_reminder_sending_response
type RenewalReminderSendingResponse struct {
	Error      *Error `json:"error,omitempty"`
	ReminderId string `json:"reminder_id,omitempty"`
	Message    string `json:"message,omitempty"`
}
