package models

import (
	"time"
)

// MFSTransaction represents a mfs_transaction
type MFSTransaction struct {
	MfsIntegrationId      string           `json:"mfs_integration_id"`
	PaymentId             string           `json:"payment_id,omitempty"`
	ProviderTransactionId string           `json:"provider_transaction_id,omitempty"`
	Amount                *Money           `json:"amount,omitempty"`
	RequestPayload        string           `json:"request_payload,omitempty"`
	TransactionId         string           `json:"transaction_id"`
	Type                  *TransactionType `json:"type"`
	CustomerMsisdn        string           `json:"customer_msisdn"`
	Status                interface{}      `json:"status"`
	ResponsePayload       string           `json:"response_payload,omitempty"`
	ErrorMessage          string           `json:"error_message,omitempty"`
	CompletedAt           time.Time        `json:"completed_at,omitempty"`
	AuditInfo             interface{}      `json:"audit_info"`
	Id                    string           `json:"id"`
	Provider              string           `json:"provider"`
}
