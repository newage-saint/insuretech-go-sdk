package models

// FieldViolation represents a field_violation
type FieldViolation struct {
	Code          string `json:"code,omitempty"`
	Description   string `json:"description,omitempty"`
	RejectedValue string `json:"rejected_value,omitempty"`
	Field         string `json:"field,omitempty"`
}
