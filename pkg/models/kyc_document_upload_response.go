package models

// KycDocumentUploadResponse represents a kyc_document_upload_response
type KycDocumentUploadResponse struct {
	Message                string `json:"message,omitempty"`
	Error                  *Error `json:"error,omitempty"`
	DocumentVerificationId string `json:"document_verification_id,omitempty"`
}
