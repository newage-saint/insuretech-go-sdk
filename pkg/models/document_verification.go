package models

// DocumentVerification represents a document_verification
type DocumentVerification struct {
	DocumentNumber    string           `json:"document_number"`
	DocumentUrl       string           `json:"document_url,omitempty"`
	Status            interface{}      `json:"status"`
	AuditInfo         interface{}      `json:"audit_info"`
	Id                string           `json:"id"`
	KycVerificationId string           `json:"kyc_verification_id"`
	ExtractedData     string           `json:"extracted_data,omitempty"`
	ConfidenceScore   float64          `json:"confidence_score,omitempty"`
	DocumentType      *KycDocumentType `json:"document_type"`
}
