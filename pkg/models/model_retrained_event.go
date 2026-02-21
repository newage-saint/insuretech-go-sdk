package models

import (
	"time"
)

// ModelRetrainedEvent represents a model_retrained_event
type ModelRetrainedEvent struct {
	OldAccuracy     float64   `json:"old_accuracy,omitempty"`
	NewAccuracy     float64   `json:"new_accuracy,omitempty"`
	TrainingSamples int       `json:"training_samples,omitempty"`
	Timestamp       time.Time `json:"timestamp,omitempty"`
	AgentId         string    `json:"agent_id,omitempty"`
	ModelName       string    `json:"model_name,omitempty"`
	CorrelationId   string    `json:"correlation_id,omitempty"`
	EventId         string    `json:"event_id,omitempty"`
	ModelVersion    string    `json:"model_version,omitempty"`
}
