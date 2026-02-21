package models

// CommissionStructureUpdateResponse represents a commission_structure_update_response
type CommissionStructureUpdateResponse struct {
	Success bool   `json:"success,omitempty"`
	Message string `json:"message,omitempty"`
	Error   *Error `json:"error,omitempty"`
}
