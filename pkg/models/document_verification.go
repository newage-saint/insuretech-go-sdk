package models

// DocumentVerification represents a document_verification
type DocumentVerification struct {
	DocumentType      *KycDocumentType `json:"document_type"`
	DocumentNumber    string           `json:"document_number"`
	DocumentUrl       string           `json:"document_url,omitempty"`
	Status            interface{}      `json:"status"`
	ConfidenceScore   float64          `json:"confidence_score,omitempty"`
	AuditInfo         *AuditInfo       `json:"audit_info,omitempty"`
	Id                string           `json:"id"`
	KycVerificationId string           `json:"kyc_verification_id"`
	ExtractedData     string           `json:"extracted_data,omitempty"`
}
