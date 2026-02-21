package models

// OTPSendingRequest represents a otp_sending_request
type OTPSendingRequest struct {
	Recipient  string `json:"recipient,omitempty"`
	Type       string `json:"type"`
	Channel    string `json:"channel,omitempty"`
	UseMasking bool   `json:"use_masking,omitempty"`
}
