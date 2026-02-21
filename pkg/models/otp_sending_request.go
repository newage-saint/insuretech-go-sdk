package models

// OTPSendingRequest represents a otp_sending_request
type OTPSendingRequest struct {
	UseMasking bool   `json:"use_masking,omitempty"`
	Recipient  string `json:"recipient,omitempty"`
	Type       string `json:"type"`
	Channel    string `json:"channel,omitempty"`
}
