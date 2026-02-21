package models

import (
	"time"
)

// FraudAlertTriggeredEvent represents a fraud_alert_triggered_event
type FraudAlertTriggeredEvent struct {
	FraudScore    int       `json:"fraud_score,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
	FraudAlertId  string    `json:"fraud_alert_id,omitempty"`
	AlertNumber   string    `json:"alert_number,omitempty"`
	EntityId      string    `json:"entity_id,omitempty"`
	CorrelationId string    `json:"correlation_id,omitempty"`
	EventId       string    `json:"event_id,omitempty"`
	EntityType    string    `json:"entity_type,omitempty"`
	RiskLevel     string    `json:"risk_level,omitempty"`
}
