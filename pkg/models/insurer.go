package models

import (
	"time"
)

// Insurer represents a insurer
type Insurer struct {
	Id                 string       `json:"id"`
	Name               string       `json:"name"`
	LogoUrl            string       `json:"logo_url,omitempty"`
	FinancialRating    string       `json:"financial_rating,omitempty"`
	PaidUpCapital      *Money       `json:"paid_up_capital,omitempty"`
	AuditInfo          interface{}  `json:"audit_info"`
	IdraLicenseExpiry  time.Time    `json:"idra_license_expiry,omitempty"`
	RegisteredAddress  interface{}  `json:"registered_address"`
	WebsiteUrl         string       `json:"website_url,omitempty"`
	Code               string       `json:"code"`
	Type               *InsurerType `json:"type"`
	Status             interface{}  `json:"status"`
	TradeLicenseNumber string       `json:"trade_license_number,omitempty"`
	IdraLicenseNumber  string       `json:"idra_license_number,omitempty"`
	ContactInfo        interface{}  `json:"contact_info"`
	HeadOfficeAddress  interface{}  `json:"head_office_address"`
	NameBn             string       `json:"name_bn,omitempty"`
	TinNumber          string       `json:"tin_number,omitempty"`
}
