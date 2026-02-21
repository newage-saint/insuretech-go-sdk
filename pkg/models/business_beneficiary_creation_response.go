package models

// BusinessBeneficiaryCreationResponse represents a business_beneficiary_creation_response
type BusinessBeneficiaryCreationResponse struct {
	BeneficiaryCode string `json:"beneficiary_code,omitempty"`
	Message         string `json:"message,omitempty"`
	Error           *Error `json:"error,omitempty"`
	BeneficiaryId   string `json:"beneficiary_id,omitempty"`
}
