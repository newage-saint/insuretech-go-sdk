package models

// FraudRule represents a fraud_rule
type FraudRule struct {
	Name        string          `json:"name"`
	Category    *RuleCategory   `json:"category"`
	Description string          `json:"description,omitempty"`
	ScoreWeight int             `json:"score_weight"`
	IsActive    bool            `json:"is_active,omitempty"`
	AuditInfo   interface{}     `json:"audit_info"`
	FraudRuleId string          `json:"fraud_rule_id"`
	Conditions  string          `json:"conditions"`
	RiskLevel   *FraudRiskLevel `json:"risk_level"`
}
