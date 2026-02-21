package models

import (
	"time"
)

// BusinessBeneficiary represents a business_beneficiary
type BusinessBeneficiary struct {
	AuditInfo              *AuditInfo      `json:"audit_info,omitempty"`
	RegistrationNumber     string          `json:"registration_number,omitempty"`
	TotalPremiumAmount     *Money          `json:"total_premium_amount,omitempty"`
	PendingActionsCount    int             `json:"pending_actions_count,omitempty"`
	Id                     string          `json:"id,omitempty"`
	BeneficiaryId          string          `json:"beneficiary_id,omitempty"`
	TradeLicenseNumber     string          `json:"trade_license_number,omitempty"`
	TinNumber              string          `json:"tin_number,omitempty"`
	FocalPersonDesignation string          `json:"focal_person_designation,omitempty"`
	PrimaryContact         *PrimaryContact `json:"primary_contact,omitempty"`
	TradeLicenseIssueDate  time.Time       `json:"trade_license_issue_date,omitempty"`
	BusinessType           *BusinessType   `json:"business_type,omitempty"`
	FocalPersonContact     *ContactInfo    `json:"focal_person_contact,omitempty"`
	TaxId                  string          `json:"tax_id,omitempty"`
	TotalEmployeesCovered  int             `json:"total_employees_covered,omitempty"`
	ActivePoliciesCount    int             `json:"active_policies_count,omitempty"`
	IncorporationDate      time.Time       `json:"incorporation_date,omitempty"`
	ContactInfo            *ContactInfo    `json:"contact_info,omitempty"`
	BusinessNameBn         string          `json:"business_name_bn,omitempty"`
	TradeLicenseExpiryDate time.Time       `json:"trade_license_expiry_date,omitempty"`
	BinNumber              string          `json:"bin_number,omitempty"`
	EmployeeCount          int             `json:"employee_count,omitempty"`
	RegisteredAddress      *Address        `json:"registered_address,omitempty"`
	BusinessAddress        *Address        `json:"business_address,omitempty"`
	BusinessName           string          `json:"business_name,omitempty"`
	IndustrySector         string          `json:"industry_sector,omitempty"`
	FocalPersonName        string          `json:"focal_person_name,omitempty"`
	FocalPersonNid         string          `json:"focal_person_nid,omitempty"`
}
