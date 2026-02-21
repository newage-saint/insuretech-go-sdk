package models

import (
	"time"
)

// ProductRiderAddedEvent represents a product_rider_added_event
type ProductRiderAddedEvent struct {
	CoverageAmount *Money    `json:"coverage_amount,omitempty"`
	IsMandatory    bool      `json:"is_mandatory,omitempty"`
	EventId        string    `json:"event_id,omitempty"`
	RiderId        string    `json:"rider_id,omitempty"`
	Timestamp      time.Time `json:"timestamp,omitempty"`
	CreatedBy      string    `json:"created_by,omitempty"`
	ProductId      string    `json:"product_id,omitempty"`
	RiderName      string    `json:"rider_name,omitempty"`
	PremiumAmount  *Money    `json:"premium_amount,omitempty"`
}
