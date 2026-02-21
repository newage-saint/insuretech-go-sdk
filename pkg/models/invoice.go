package models

import (
	"time"
)

// Invoice represents a invoice
type Invoice struct {
	Amount        *Money         `json:"amount,omitempty"`
	DueDate       time.Time      `json:"due_date,omitempty"`
	Status        *InvoiceStatus `json:"status,omitempty"`
	PaymentId     string         `json:"payment_id,omitempty"`
	PolicyIds     []string       `json:"policy_ids,omitempty"`
	CreatedAt     time.Time      `json:"created_at,omitempty"`
	InvoiceId     string         `json:"invoice_id,omitempty"`
	IssuedAt      time.Time      `json:"issued_at,omitempty"`
	PaidAt        time.Time      `json:"paid_at,omitempty"`
	InvoicePdfUrl string         `json:"invoice_pdf_url,omitempty"`
	UpdatedAt     time.Time      `json:"updated_at,omitempty"`
	InvoiceNumber string         `json:"invoice_number,omitempty"`
	BusinessId    string         `json:"business_id,omitempty"`
}
