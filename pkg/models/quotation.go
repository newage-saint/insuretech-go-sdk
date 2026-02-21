package models

import (
	"time"
)

// Quotation represents a quotation
type Quotation struct {
	ApprovedAt        time.Time        `json:"approved_at,omitempty"`
	BusinessId        string           `json:"business_id,omitempty"`
	InsurerName       string           `json:"insurer_name,omitempty"`
	EstimatedPremium  *Money           `json:"estimated_premium,omitempty"`
	RejectionReason   string           `json:"rejection_reason,omitempty"`
	InsuranceCategory *InsuranceType   `json:"insurance_category,omitempty"`
	EmployeeNo        int              `json:"employee_no,omitempty"`
	Status            *QuotationStatus `json:"status,omitempty"`
	CreatedAt         time.Time        `json:"created_at,omitempty"`
	UpdatedAt         time.Time        `json:"updated_at,omitempty"`
	QuotationNumber   string           `json:"quotation_number,omitempty"`
	PlanName          string           `json:"plan_name,omitempty"`
	CreatedByUserId   string           `json:"created_by_user_id,omitempty"`
	QuotationId       string           `json:"quotation_id,omitempty"`
	PlanId            string           `json:"plan_id,omitempty"`
	DepartmentId      string           `json:"department_id,omitempty"`
	QuotedAmount      *Money           `json:"quoted_amount,omitempty"`
	SubmissionDate    time.Time        `json:"submission_date,omitempty"`
	ValidUntil        time.Time        `json:"valid_until,omitempty"`
	ApprovedByUserId  string           `json:"approved_by_user_id,omitempty"`
}
