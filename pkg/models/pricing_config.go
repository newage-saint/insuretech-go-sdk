package models

import (
	"time"
)

// PricingConfig represents a pricing_config
type PricingConfig struct {
	EffectiveFrom   time.Time      `json:"effective_from"`
	EffectiveTo     time.Time      `json:"effective_to,omitempty"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	PricingConfigId string         `json:"pricing_config_id"`
	ProductId       string         `json:"product_id"`
	Rules           []*PricingRule `json:"rules,omitempty"`
}
