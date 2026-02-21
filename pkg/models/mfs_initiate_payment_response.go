package models

// MfsInitiatePaymentResponse represents a mfs_initiate_payment_response
type MfsInitiatePaymentResponse struct {
	Message          string `json:"message,omitempty"`
	Error            *Error `json:"error,omitempty"`
	MfsTransactionId string `json:"mfs_transaction_id,omitempty"`
	TransactionId    string `json:"transaction_id,omitempty"`
	PaymentUrl       string `json:"payment_url,omitempty"`
}
