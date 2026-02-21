package models

// RuleCondition represents a rule_condition
type RuleCondition struct {
	Operator string `json:"operator,omitempty"`
	Value    string `json:"value,omitempty"`
	Field    string `json:"field,omitempty"`
}
