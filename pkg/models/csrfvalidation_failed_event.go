package models

import (
	"time"
)

// CSRFValidationFailedEvent represents a csrfvalidation_failed_event
type CSRFValidationFailedEvent struct {
	RequestPath           string    `json:"request_path,omitempty"`
	RequestMethod         string    `json:"request_method,omitempty"`
	EventId               string    `json:"event_id,omitempty"`
	UserId                string    `json:"user_id,omitempty"`
	ReceivedCsrfTokenHash string    `json:"received_csrf_token_hash,omitempty"`
	Timestamp             time.Time `json:"timestamp,omitempty"`
	IpAddress             string    `json:"ip_address,omitempty"`
	SessionId             string    `json:"session_id,omitempty"`
	ExpectedCsrfTokenHash string    `json:"expected_csrf_token_hash,omitempty"`
	UserAgent             string    `json:"user_agent,omitempty"`
}
