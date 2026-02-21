package models

// RuleCondition represents a rule_condition
type RuleCondition struct {
	Value    string `json:"value,omitempty"`
	Field    string `json:"field,omitempty"`
	Operator string `json:"operator,omitempty"`
}
