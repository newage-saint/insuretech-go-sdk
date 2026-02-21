package models

// DocumentVerification represents a document_verification
type DocumentVerification struct {
	DocumentType      *KycDocumentType `json:"document_type"`
	DocumentNumber    string           `json:"document_number"`
	ExtractedData     string           `json:"extracted_data,omitempty"`
	Status            interface{}      `json:"status"`
	Id                string           `json:"id"`
	KycVerificationId string           `json:"kyc_verification_id"`
	DocumentUrl       string           `json:"document_url,omitempty"`
	ConfidenceScore   float64          `json:"confidence_score,omitempty"`
	AuditInfo         interface{}      `json:"audit_info"`
}
