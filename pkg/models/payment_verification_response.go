package models

// PaymentVerificationResponse represents a payment_verification_response
type PaymentVerificationResponse struct {
	Verified  bool     `json:"verified,omitempty"`
	Error     *Error   `json:"error,omitempty"`
	PaymentId string   `json:"payment_id,omitempty"`
	Status    string   `json:"status,omitempty"`
	Payment   *Payment `json:"payment,omitempty"`
}
