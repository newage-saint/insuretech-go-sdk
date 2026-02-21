package models

import (
	"time"
)

// ProductPlan represents a product_plan
type ProductPlan struct {
	PlanId          string    `json:"plan_id,omitempty"`
	PlanName        string    `json:"plan_name,omitempty"`
	PremiumAmount   *Money    `json:"premium_amount,omitempty"`
	MinSumInsured   *Money    `json:"min_sum_insured,omitempty"`
	MaxSumInsured   *Money    `json:"max_sum_insured,omitempty"`
	Attributes      string    `json:"attributes,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
	ProductId       string    `json:"product_id,omitempty"`
	PlanDescription string    `json:"plan_description,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
}
