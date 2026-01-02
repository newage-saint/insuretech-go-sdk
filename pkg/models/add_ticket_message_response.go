package models

// AddTicketMessageResponse represents a add_ticket_message_response
type AddTicketMessageResponse struct {
	Error     *Error `json:"error,omitempty"`
	MessageId string `json:"message_id,omitempty"`
	Message   string `json:"message,omitempty"`
}
