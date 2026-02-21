package models

// SettleClaimResponse represents a settle_claim_response
type SettleClaimResponse struct {
	Error         *Error `json:"error,omitempty"`
	Message       string `json:"message,omitempty"`
	SettledAmount *Money `json:"settled_amount,omitempty"`
	PaymentId     string `json:"payment_id,omitempty"`
}
