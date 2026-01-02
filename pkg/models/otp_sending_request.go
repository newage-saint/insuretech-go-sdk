package models

// OTPSendingRequest represents a otp_sending_request
type OTPSendingRequest struct {
	Type      string `json:"type"`
	Recipient string `json:"recipient,omitempty"`
}
