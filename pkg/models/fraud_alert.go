package models

import (
	"time"
)

// FraudAlert represents a fraud_alert
type FraudAlert struct {
	Id          string      `json:"id"`
	AlertNumber string      `json:"alert_number"`
	EntityId    string      `json:"entity_id"`
	FraudRuleId string      `json:"fraud_rule_id"`
	RiskLevel   string      `json:"risk_level"`
	FraudScore  int         `json:"fraud_score"`
	Details     string      `json:"details,omitempty"`
	Status      interface{} `json:"status"`
	EntityType  string      `json:"entity_type"`
	AssignedTo  string      `json:"assigned_to,omitempty"`
	ResolvedAt  time.Time   `json:"resolved_at,omitempty"`
	AuditInfo   interface{} `json:"audit_info"`
}
