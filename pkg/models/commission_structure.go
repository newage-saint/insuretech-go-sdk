package models

// CommissionStructure represents a commission_structure
type CommissionStructure struct {
	RenewalRate          float64 `json:"renewal_rate,omitempty"`
	ClaimsAssistanceRate float64 `json:"claims_assistance_rate,omitempty"`
	AcquisitionRate      float64 `json:"acquisition_rate,omitempty"`
}
