package models

// Condition represents a condition
type Condition struct {
	Values    []string  `json:"values,omitempty"`
	Attribute string    `json:"attribute,omitempty"`
	Operator  *Operator `json:"operator,omitempty"`
}
