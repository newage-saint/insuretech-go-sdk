package models

import (
	"time"
)

// MFSTransaction represents a mfs_transaction
type MFSTransaction struct {
	Provider              string           `json:"provider"`
	Amount                *Money           `json:"amount,omitempty"`
	CustomerMsisdn        string           `json:"customer_msisdn"`
	RequestPayload        string           `json:"request_payload,omitempty"`
	CompletedAt           time.Time        `json:"completed_at,omitempty"`
	TransactionId         string           `json:"transaction_id"`
	PaymentId             string           `json:"payment_id,omitempty"`
	Type                  *TransactionType `json:"type"`
	ProviderTransactionId string           `json:"provider_transaction_id,omitempty"`
	ErrorMessage          string           `json:"error_message,omitempty"`
	Id                    string           `json:"id"`
	Status                interface{}      `json:"status"`
	ResponsePayload       string           `json:"response_payload,omitempty"`
	AuditInfo             interface{}      `json:"audit_info"`
	MfsIntegrationId      string           `json:"mfs_integration_id"`
}
