package models

import (
	"time"
)

// InsurerProduct represents a insurer_product
type InsurerProduct struct {
	Id                  string      `json:"id"`
	MinSumAssured       *Money      `json:"min_sum_assured,omitempty"`
	MinEntryAge         int         `json:"min_entry_age,omitempty"`
	MinTermYears        int         `json:"min_term_years,omitempty"`
	MedicalRequired     bool        `json:"medical_required,omitempty"`
	MedicalThreshold    *Money      `json:"medical_threshold,omitempty"`
	FreeLookPeriodDays  int         `json:"free_look_period_days,omitempty"`
	CommissionConfigId  string      `json:"commission_config_id,omitempty"`
	InsurerId           string      `json:"insurer_id"`
	Name                string      `json:"name"`
	MaxSumAssured       *Money      `json:"max_sum_assured,omitempty"`
	MaxEntryAge         int         `json:"max_entry_age,omitempty"`
	MaxTermYears        int         `json:"max_term_years,omitempty"`
	Exclusions          string      `json:"exclusions,omitempty"`
	EffectiveFrom       time.Time   `json:"effective_from"`
	EffectiveTo         time.Time   `json:"effective_to,omitempty"`
	ProductId           string      `json:"product_id"`
	Code                string      `json:"code"`
	MaxPremium          *Money      `json:"max_premium,omitempty"`
	MaxMaturityAge      int         `json:"max_maturity_age,omitempty"`
	PremiumPaymentModes []string    `json:"premium_payment_modes,omitempty"`
	Status              interface{} `json:"status"`
	MinPremium          *Money      `json:"min_premium,omitempty"`
	Features            string      `json:"features,omitempty"`
	AuditInfo           interface{} `json:"audit_info"`
}
