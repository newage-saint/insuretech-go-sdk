package models

// PaymentVerificationResponse represents a payment_verification_response
type PaymentVerificationResponse struct {
	Error     *Error   `json:"error,omitempty"`
	PaymentId string   `json:"payment_id,omitempty"`
	Status    string   `json:"status,omitempty"`
	Payment   *Payment `json:"payment,omitempty"`
	Verified  bool     `json:"verified,omitempty"`
}
