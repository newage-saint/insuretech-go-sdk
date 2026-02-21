package models

import (
	"time"
)

// Partner represents a partner
type Partner struct {
	BankBranch                string               `json:"bank_branch,omitempty"`
	RenewalCommissionRate     float64              `json:"renewal_commission_rate"`
	OnboardedAt               time.Time            `json:"onboarded_at,omitempty"`
	PartnerId                 string               `json:"partner_id"`
	OrganizationName          string               `json:"organization_name"`
	BankName                  string               `json:"bank_name,omitempty"`
	ContactEmail              string               `json:"contact_email"`
	ContactPhone              string               `json:"contact_phone"`
	UpdatedAt                 time.Time            `json:"updated_at"`
	DeletedAt                 time.Time            `json:"deleted_at,omitempty"`
	Status                    interface{}          `json:"status"`
	TinNumber                 string               `json:"tin_number"`
	BankAccount               string               `json:"bank_account"`
	CreatedAt                 time.Time            `json:"created_at"`
	Type                      *PartnerType         `json:"type"`
	TradeLicense              string               `json:"trade_license"`
	AcquisitionCommissionRate float64              `json:"acquisition_commission_rate"`
	ClaimsAssistanceRate      float64              `json:"claims_assistance_rate,omitempty"`
	FocalPersonId             string               `json:"focal_person_id,omitempty"`
	Commission                *CommissionStructure `json:"commission,omitempty"`
	Benefits                  *PartnerBenefits     `json:"benefits,omitempty"`
}
