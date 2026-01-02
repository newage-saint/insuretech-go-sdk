package models

// TicketAssignmentResponse represents a ticket_assignment_response
type TicketAssignmentResponse struct {
	Error   *Error `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}
