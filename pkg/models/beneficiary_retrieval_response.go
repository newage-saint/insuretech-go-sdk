package models

// BeneficiaryRetrievalResponse represents a beneficiary_retrieval_response
type BeneficiaryRetrievalResponse struct {
	BusinessDetails   *BusinessBeneficiary   `json:"business_details,omitempty"`
	Error             *Error                 `json:"error,omitempty"`
	Beneficiary       *Beneficiary           `json:"beneficiary,omitempty"`
	IndividualDetails *IndividualBeneficiary `json:"individual_details,omitempty"`
}
