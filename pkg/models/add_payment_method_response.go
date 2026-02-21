package models

// AddPaymentMethodResponse represents a add_payment_method_response
type AddPaymentMethodResponse struct {
	Success  bool   `json:"success,omitempty"`
	Error    *Error `json:"error,omitempty"`
	MethodId string `json:"method_id,omitempty"`
}
