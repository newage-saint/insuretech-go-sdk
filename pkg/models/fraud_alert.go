package models

import (
	"time"
)

// FraudAlert represents a fraud_alert
type FraudAlert struct {
	AssignedTo  string      `json:"assigned_to,omitempty"`
	ResolvedAt  time.Time   `json:"resolved_at,omitempty"`
	AuditInfo   interface{} `json:"audit_info"`
	EntityType  string      `json:"entity_type"`
	EntityId    string      `json:"entity_id"`
	Details     string      `json:"details,omitempty"`
	Status      interface{} `json:"status"`
	Id          string      `json:"id"`
	AlertNumber string      `json:"alert_number"`
	FraudRuleId string      `json:"fraud_rule_id"`
	RiskLevel   string      `json:"risk_level"`
	FraudScore  int         `json:"fraud_score"`
}
