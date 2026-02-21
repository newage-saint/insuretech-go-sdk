package models

import (
	"time"
)

// Quotation represents a quotation
type Quotation struct {
	InsuranceCategory *InsuranceType   `json:"insurance_category,omitempty"`
	EmployeeNo        int              `json:"employee_no,omitempty"`
	QuotedAmount      *Money           `json:"quoted_amount,omitempty"`
	ApprovedAt        time.Time        `json:"approved_at,omitempty"`
	RejectionReason   string           `json:"rejection_reason,omitempty"`
	ValidUntil        time.Time        `json:"valid_until,omitempty"`
	CreatedAt         time.Time        `json:"created_at,omitempty"`
	CreatedByUserId   string           `json:"created_by_user_id,omitempty"`
	BusinessId        string           `json:"business_id,omitempty"`
	InsurerName       string           `json:"insurer_name,omitempty"`
	DepartmentId      string           `json:"department_id,omitempty"`
	Status            *QuotationStatus `json:"status,omitempty"`
	SubmissionDate    time.Time        `json:"submission_date,omitempty"`
	PlanName          string           `json:"plan_name,omitempty"`
	ApprovedByUserId  string           `json:"approved_by_user_id,omitempty"`
	QuotationId       string           `json:"quotation_id,omitempty"`
	PlanId            string           `json:"plan_id,omitempty"`
	EstimatedPremium  *Money           `json:"estimated_premium,omitempty"`
	UpdatedAt         time.Time        `json:"updated_at,omitempty"`
	QuotationNumber   string           `json:"quotation_number,omitempty"`
}
