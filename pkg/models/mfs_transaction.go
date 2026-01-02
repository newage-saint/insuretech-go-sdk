package models

import (
	"time"
)

// MFSTransaction represents a mfs_transaction
type MFSTransaction struct {
	PaymentId             string           `json:"payment_id,omitempty"`
	Type                  *TransactionType `json:"type"`
	Status                interface{}      `json:"status"`
	RequestPayload        string           `json:"request_payload,omitempty"`
	Amount                *Money           `json:"amount,omitempty"`
	CustomerMsisdn        string           `json:"customer_msisdn"`
	ErrorMessage          string           `json:"error_message,omitempty"`
	Id                    string           `json:"id"`
	TransactionId         string           `json:"transaction_id"`
	ProviderTransactionId string           `json:"provider_transaction_id,omitempty"`
	AuditInfo             *AuditInfo       `json:"audit_info,omitempty"`
	MfsIntegrationId      string           `json:"mfs_integration_id"`
	ResponsePayload       string           `json:"response_payload,omitempty"`
	CompletedAt           time.Time        `json:"completed_at,omitempty"`
	Provider              string           `json:"provider"`
}
