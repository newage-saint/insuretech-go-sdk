package models

// PaymentStatus represents a payment_status
type PaymentStatus string

// PaymentStatus values
const (
	PaymentStatusPAYMENTSTATUSUNSPECIFIED PaymentStatus = "PAYMENT_STATUS_UNSPECIFIED"
	PaymentStatusPAYMENTSTATUSINITIATED                 = "PAYMENT_STATUS_INITIATED"
	PaymentStatusPAYMENTSTATUSPENDING                   = "PAYMENT_STATUS_PENDING"
	PaymentStatusPAYMENTSTATUSPROCESSING                = "PAYMENT_STATUS_PROCESSING"
	PaymentStatusPAYMENTSTATUSSUCCESS                   = "PAYMENT_STATUS_SUCCESS"
	PaymentStatusPAYMENTSTATUSFAILED                    = "PAYMENT_STATUS_FAILED"
	PaymentStatusPAYMENTSTATUSREFUNDED                  = "PAYMENT_STATUS_REFUNDED"
	PaymentStatusPAYMENTSTATUSCANCELLED                 = "PAYMENT_STATUS_CANCELLED"
)
