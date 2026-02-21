package models

import (
	"time"
)

// Insurer represents a insurer
type Insurer struct {
	LogoUrl            string       `json:"logo_url,omitempty"`
	Code               string       `json:"code"`
	NameBn             string       `json:"name_bn,omitempty"`
	IdraLicenseNumber  string       `json:"idra_license_number,omitempty"`
	WebsiteUrl         string       `json:"website_url,omitempty"`
	PaidUpCapital      *Money       `json:"paid_up_capital,omitempty"`
	Status             interface{}  `json:"status"`
	RegisteredAddress  interface{}  `json:"registered_address"`
	Id                 string       `json:"id"`
	IdraLicenseExpiry  time.Time    `json:"idra_license_expiry,omitempty"`
	ContactInfo        interface{}  `json:"contact_info"`
	HeadOfficeAddress  interface{}  `json:"head_office_address"`
	FinancialRating    string       `json:"financial_rating,omitempty"`
	AuditInfo          interface{}  `json:"audit_info"`
	Name               string       `json:"name"`
	Type               *InsurerType `json:"type"`
	TradeLicenseNumber string       `json:"trade_license_number,omitempty"`
	TinNumber          string       `json:"tin_number,omitempty"`
}
