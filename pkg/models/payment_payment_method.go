package models

// PaymentPaymentMethod represents a payment_payment_method
type PaymentPaymentMethod struct {
	DisplayName string                 `json:"display_name,omitempty"`
	IsDefault   bool                   `json:"is_default,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	MethodId    string                 `json:"method_id,omitempty"`
	MethodType  string                 `json:"method_type,omitempty"`
}
