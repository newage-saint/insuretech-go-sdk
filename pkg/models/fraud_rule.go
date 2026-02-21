package models

// FraudRule represents a fraud_rule
type FraudRule struct {
	Category    *RuleCategory   `json:"category"`
	Description string          `json:"description,omitempty"`
	ScoreWeight int             `json:"score_weight"`
	AuditInfo   interface{}     `json:"audit_info"`
	FraudRuleId string          `json:"fraud_rule_id"`
	Name        string          `json:"name"`
	Conditions  string          `json:"conditions"`
	RiskLevel   *FraudRiskLevel `json:"risk_level"`
	IsActive    bool            `json:"is_active,omitempty"`
}
