package models

import (
	"time"
)

// Product represents a product
type Product struct {
	DeletedAt       time.Time        `json:"deleted_at,omitempty"`
	ProductCode     string           `json:"product_code"`
	MinTenureMonths int              `json:"min_tenure_months"`
	CreatedAt       time.Time        `json:"created_at"`
	Category        *ProductCategory `json:"category"`
	Exclusions      []string         `json:"exclusions,omitempty"`
	UpdatedAt       time.Time        `json:"updated_at"`
	CreatedBy       string           `json:"created_by"`
	PricingConfig   *PricingConfig   `json:"pricing_config,omitempty"`
	Description     string           `json:"description,omitempty"`
	MaxTenureMonths int              `json:"max_tenure_months"`
	AvailableRiders []*ProductsRider `json:"available_riders,omitempty"`
	ProductId       string           `json:"product_id"`
	ProductName     string           `json:"product_name"`
	BasePremium     *Money           `json:"base_premium"`
	MinSumInsured   *Money           `json:"min_sum_insured"`
	MaxSumInsured   *Money           `json:"max_sum_insured"`
	Status          interface{}      `json:"status"`
}
