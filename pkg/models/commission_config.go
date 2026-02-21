package models

import (
	"time"
)

// CommissionConfig represents a commission_config
type CommissionConfig struct {
	AgentSplitConfig     string        `json:"agent_split_config,omitempty"`
	EffectiveFrom        time.Time     `json:"effective_from"`
	Id                   string        `json:"id"`
	RevenueModel         *RevenueModel `json:"revenue_model"`
	AcquisitionFlatFee   *Money        `json:"acquisition_flat_fee,omitempty"`
	ClaimsAssistanceRate string        `json:"claims_assistance_rate,omitempty"`
	PerformanceTiers     string        `json:"performance_tiers,omitempty"`
	EffectiveTo          time.Time     `json:"effective_to,omitempty"`
	AuditInfo            interface{}   `json:"audit_info"`
	InsurerProductId     string        `json:"insurer_product_id"`
	AcquisitionRate      string        `json:"acquisition_rate,omitempty"`
	RenewalRate          string        `json:"renewal_rate,omitempty"`
	RenewalFlatFee       *Money        `json:"renewal_flat_fee,omitempty"`
}
