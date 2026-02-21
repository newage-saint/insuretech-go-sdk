package models

import (
	"time"
)

// FraudAlertTriggeredEvent represents a fraud_alert_triggered_event
type FraudAlertTriggeredEvent struct {
	EventId       string    `json:"event_id,omitempty"`
	FraudAlertId  string    `json:"fraud_alert_id,omitempty"`
	AlertNumber   string    `json:"alert_number,omitempty"`
	RiskLevel     string    `json:"risk_level,omitempty"`
	FraudScore    int       `json:"fraud_score,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
	EntityType    string    `json:"entity_type,omitempty"`
	EntityId      string    `json:"entity_id,omitempty"`
	CorrelationId string    `json:"correlation_id,omitempty"`
}
