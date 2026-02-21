package models

import (
	"time"
)

// Insurer represents a insurer
type Insurer struct {
	WebsiteUrl         string       `json:"website_url,omitempty"`
	PaidUpCapital      *Money       `json:"paid_up_capital,omitempty"`
	Id                 string       `json:"id"`
	ContactInfo        interface{}  `json:"contact_info"`
	Name               string       `json:"name"`
	Type               *InsurerType `json:"type"`
	TradeLicenseNumber string       `json:"trade_license_number,omitempty"`
	RegisteredAddress  interface{}  `json:"registered_address"`
	HeadOfficeAddress  interface{}  `json:"head_office_address"`
	LogoUrl            string       `json:"logo_url,omitempty"`
	FinancialRating    string       `json:"financial_rating,omitempty"`
	AuditInfo          interface{}  `json:"audit_info"`
	Code               string       `json:"code"`
	NameBn             string       `json:"name_bn,omitempty"`
	Status             interface{}  `json:"status"`
	TinNumber          string       `json:"tin_number,omitempty"`
	IdraLicenseNumber  string       `json:"idra_license_number,omitempty"`
	IdraLicenseExpiry  time.Time    `json:"idra_license_expiry,omitempty"`
}
