package models

import (
	"time"
)

// CommissionConfig represents a commission_config
type CommissionConfig struct {
	ClaimsAssistanceRate string        `json:"claims_assistance_rate,omitempty"`
	AgentSplitConfig     string        `json:"agent_split_config,omitempty"`
	PerformanceTiers     string        `json:"performance_tiers,omitempty"`
	EffectiveFrom        time.Time     `json:"effective_from"`
	EffectiveTo          time.Time     `json:"effective_to,omitempty"`
	Id                   string        `json:"id"`
	InsurerProductId     string        `json:"insurer_product_id"`
	AcquisitionFlatFee   *Money        `json:"acquisition_flat_fee,omitempty"`
	RenewalFlatFee       *Money        `json:"renewal_flat_fee,omitempty"`
	AuditInfo            interface{}   `json:"audit_info"`
	RevenueModel         *RevenueModel `json:"revenue_model"`
	AcquisitionRate      string        `json:"acquisition_rate,omitempty"`
	RenewalRate          string        `json:"renewal_rate,omitempty"`
}
