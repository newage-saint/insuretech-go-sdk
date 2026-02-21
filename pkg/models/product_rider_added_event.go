package models

import (
	"time"
)

// ProductRiderAddedEvent represents a product_rider_added_event
type ProductRiderAddedEvent struct {
	IsMandatory    bool      `json:"is_mandatory,omitempty"`
	EventId        string    `json:"event_id,omitempty"`
	RiderName      string    `json:"rider_name,omitempty"`
	CoverageAmount *Money    `json:"coverage_amount,omitempty"`
	Timestamp      time.Time `json:"timestamp,omitempty"`
	CreatedBy      string    `json:"created_by,omitempty"`
	RiderId        string    `json:"rider_id,omitempty"`
	ProductId      string    `json:"product_id,omitempty"`
	PremiumAmount  *Money    `json:"premium_amount,omitempty"`
}
