package models

import (
	"time"
)

// MFSTransaction represents a mfs_transaction
type MFSTransaction struct {
	Provider              string           `json:"provider"`
	ProviderTransactionId string           `json:"provider_transaction_id,omitempty"`
	PaymentId             string           `json:"payment_id,omitempty"`
	Type                  *TransactionType `json:"type"`
	Status                interface{}      `json:"status"`
	ResponsePayload       string           `json:"response_payload,omitempty"`
	ErrorMessage          string           `json:"error_message,omitempty"`
	AuditInfo             interface{}      `json:"audit_info"`
	TransactionId         string           `json:"transaction_id"`
	MfsIntegrationId      string           `json:"mfs_integration_id"`
	Amount                *Money           `json:"amount,omitempty"`
	CustomerMsisdn        string           `json:"customer_msisdn"`
	Id                    string           `json:"id"`
	RequestPayload        string           `json:"request_payload,omitempty"`
	CompletedAt           time.Time        `json:"completed_at,omitempty"`
}
