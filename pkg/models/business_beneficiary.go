package models

import (
	"time"
)

// BusinessBeneficiary represents a business_beneficiary
type BusinessBeneficiary struct {
	TotalEmployeesCovered  int             `json:"total_employees_covered,omitempty"`
	BeneficiaryId          string          `json:"beneficiary_id,omitempty"`
	TradeLicenseNumber     string          `json:"trade_license_number,omitempty"`
	TradeLicenseIssueDate  time.Time       `json:"trade_license_issue_date,omitempty"`
	AuditInfo              *AuditInfo      `json:"audit_info,omitempty"`
	ActivePoliciesCount    int             `json:"active_policies_count,omitempty"`
	TotalPremiumAmount     *Money          `json:"total_premium_amount,omitempty"`
	PendingActionsCount    int             `json:"pending_actions_count,omitempty"`
	ContactInfo            *ContactInfo    `json:"contact_info,omitempty"`
	Id                     string          `json:"id,omitempty"`
	BusinessName           string          `json:"business_name,omitempty"`
	TradeLicenseExpiryDate time.Time       `json:"trade_license_expiry_date,omitempty"`
	IndustrySector         string          `json:"industry_sector,omitempty"`
	EmployeeCount          int             `json:"employee_count,omitempty"`
	IncorporationDate      time.Time       `json:"incorporation_date,omitempty"`
	RegisteredAddress      *Address        `json:"registered_address,omitempty"`
	BusinessNameBn         string          `json:"business_name_bn,omitempty"`
	TinNumber              string          `json:"tin_number,omitempty"`
	FocalPersonName        string          `json:"focal_person_name,omitempty"`
	FocalPersonNid         string          `json:"focal_person_nid,omitempty"`
	FocalPersonContact     *ContactInfo    `json:"focal_person_contact,omitempty"`
	RegistrationNumber     string          `json:"registration_number,omitempty"`
	TaxId                  string          `json:"tax_id,omitempty"`
	PrimaryContact         *PrimaryContact `json:"primary_contact,omitempty"`
	BinNumber              string          `json:"bin_number,omitempty"`
	BusinessType           *BusinessType   `json:"business_type,omitempty"`
	BusinessAddress        *Address        `json:"business_address,omitempty"`
	FocalPersonDesignation string          `json:"focal_person_designation,omitempty"`
}
