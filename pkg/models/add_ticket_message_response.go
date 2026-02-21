package models

// AddTicketMessageResponse represents a add_ticket_message_response
type AddTicketMessageResponse struct {
	MessageId string `json:"message_id,omitempty"`
	Message   string `json:"message,omitempty"`
	Error     *Error `json:"error,omitempty"`
}
