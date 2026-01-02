package models

// PremiumBreakdown represents a premium_breakdown
type PremiumBreakdown struct {
	Item        string `json:"item,omitempty"`
	Amount      *Money `json:"amount,omitempty"`
	Description string `json:"description,omitempty"`
}
