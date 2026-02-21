package models

// DocumentVerification represents a document_verification
type DocumentVerification struct {
	Id                string           `json:"id"`
	KycVerificationId string           `json:"kyc_verification_id"`
	DocumentUrl       string           `json:"document_url,omitempty"`
	ExtractedData     string           `json:"extracted_data,omitempty"`
	Status            interface{}      `json:"status"`
	ConfidenceScore   float64          `json:"confidence_score,omitempty"`
	AuditInfo         interface{}      `json:"audit_info"`
	DocumentType      *KycDocumentType `json:"document_type"`
	DocumentNumber    string           `json:"document_number"`
}
