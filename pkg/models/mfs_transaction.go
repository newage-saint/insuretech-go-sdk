package models

import (
	"time"
)

// MFSTransaction represents a mfs_transaction
type MFSTransaction struct {
	Amount                *Money           `json:"amount,omitempty"`
	CustomerMsisdn        string           `json:"customer_msisdn"`
	RequestPayload        string           `json:"request_payload,omitempty"`
	CompletedAt           time.Time        `json:"completed_at,omitempty"`
	AuditInfo             interface{}      `json:"audit_info"`
	Id                    string           `json:"id"`
	TransactionId         string           `json:"transaction_id"`
	ProviderTransactionId string           `json:"provider_transaction_id,omitempty"`
	Status                interface{}      `json:"status"`
	MfsIntegrationId      string           `json:"mfs_integration_id"`
	Type                  *TransactionType `json:"type"`
	ErrorMessage          string           `json:"error_message,omitempty"`
	Provider              string           `json:"provider"`
	PaymentId             string           `json:"payment_id,omitempty"`
	ResponsePayload       string           `json:"response_payload,omitempty"`
}
