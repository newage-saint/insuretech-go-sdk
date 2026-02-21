package models

// PaymentInitiatePaymentRequest represents a payment_initiate_payment_request
type PaymentInitiatePaymentRequest struct {
	UserId         string                 `json:"user_id"`
	PolicyId       string                 `json:"policy_id"`
	Amount         *Money                 `json:"amount,omitempty"`
	Currency       string                 `json:"currency,omitempty"`
	PaymentMethod  string                 `json:"payment_method,omitempty"`
	CallbackUrl    string                 `json:"callback_url,omitempty"`
	Metadata       map[string]interface{} `json:"metadata,omitempty"`
	IdempotencyKey string                 `json:"idempotency_key,omitempty"`
}
