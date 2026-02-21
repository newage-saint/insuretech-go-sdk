package models

// InvoiceStatus represents a invoice_status
type InvoiceStatus string

// InvoiceStatus values
const (
	InvoiceStatusINVOICESTATUSUNSPECIFIED InvoiceStatus = "INVOICE_STATUS_UNSPECIFIED"
	InvoiceStatusINVOICESTATUSPENDING                   = "INVOICE_STATUS_PENDING"
	InvoiceStatusINVOICESTATUSAPPROVED                  = "INVOICE_STATUS_APPROVED"
	InvoiceStatusINVOICESTATUSPAID                      = "INVOICE_STATUS_PAID"
	InvoiceStatusINVOICESTATUSOVERDUE                   = "INVOICE_STATUS_OVERDUE"
	InvoiceStatusINVOICESTATUSCANCELLED                 = "INVOICE_STATUS_CANCELLED"
)
