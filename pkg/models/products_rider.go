package models

import (
	"time"
)

// ProductsRider represents a products_rider
type ProductsRider struct {
	ProductId      string    `json:"product_id"`
	RiderName      string    `json:"rider_name"`
	Description    string    `json:"description,omitempty"`
	PremiumAmount  *Money    `json:"premium_amount"`
	IsMandatory    bool      `json:"is_mandatory"`
	CreatedAt      time.Time `json:"created_at"`
	RiderId        string    `json:"rider_id"`
	CoverageAmount *Money    `json:"coverage_amount"`
	UpdatedAt      time.Time `json:"updated_at"`
}
