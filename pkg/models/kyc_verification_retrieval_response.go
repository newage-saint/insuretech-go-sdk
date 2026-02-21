package models

// KYCVerificationRetrievalResponse represents a kyc_verification_retrieval_response
type KYCVerificationRetrievalResponse struct {
	KycVerification *KYCVerification        `json:"kyc_verification,omitempty"`
	Documents       []*DocumentVerification `json:"documents,omitempty"`
	Error           *Error                  `json:"error,omitempty"`
}
