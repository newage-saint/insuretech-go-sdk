package models

import (
	"time"
)

// Product represents a product
type Product struct {
	MinSumInsured         *Money                 `json:"min_sum_insured,omitempty"`
	UpdatedAt             time.Time              `json:"updated_at,omitempty"`
	CreatedBy             string                 `json:"created_by,omitempty"`
	PricingConfig         *PricingConfig         `json:"pricing_config,omitempty"`
	ProductCode           string                 `json:"product_code,omitempty"`
	ProductName           string                 `json:"product_name,omitempty"`
	Category              *ProductCategory       `json:"category,omitempty"`
	BasePremium           *Money                 `json:"base_premium,omitempty"`
	MaxTenureMonths       int                    `json:"max_tenure_months,omitempty"`
	Exclusions            []string               `json:"exclusions,omitempty"`
	MaxSumInsuredCurrency string                 `json:"max_sum_insured_currency,omitempty"`
	MaxSumInsured         *Money                 `json:"max_sum_insured,omitempty"`
	MinTenureMonths       int                    `json:"min_tenure_months,omitempty"`
	CreatedAt             time.Time              `json:"created_at,omitempty"`
	DeletedAt             time.Time              `json:"deleted_at,omitempty"`
	AvailableRiders       []*ProductsRider       `json:"available_riders,omitempty"`
	ProductAttributes     string                 `json:"product_attributes,omitempty"`
	BasePremiumCurrency   string                 `json:"base_premium_currency,omitempty"`
	ProductId             string                 `json:"product_id,omitempty"`
	Description           string                 `json:"description,omitempty"`
	Status                *ProductsProductStatus `json:"status,omitempty"`
	Plans                 []*ProductPlan         `json:"plans,omitempty"`
	MinSumInsuredCurrency string                 `json:"min_sum_insured_currency,omitempty"`
}
