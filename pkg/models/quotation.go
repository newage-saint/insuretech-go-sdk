package models

import (
	"time"
)

// Quotation represents a quotation
type Quotation struct {
	QuotedAmount      *Money           `json:"quoted_amount,omitempty"`
	Status            *QuotationStatus `json:"status,omitempty"`
	UpdatedAt         time.Time        `json:"updated_at,omitempty"`
	QuotationId       string           `json:"quotation_id,omitempty"`
	BusinessId        string           `json:"business_id,omitempty"`
	EstimatedPremium  *Money           `json:"estimated_premium,omitempty"`
	ApprovedByUserId  string           `json:"approved_by_user_id,omitempty"`
	InsuranceCategory *InsuranceType   `json:"insurance_category,omitempty"`
	EmployeeNo        int              `json:"employee_no,omitempty"`
	CreatedAt         time.Time        `json:"created_at,omitempty"`
	ApprovedAt        time.Time        `json:"approved_at,omitempty"`
	RejectionReason   string           `json:"rejection_reason,omitempty"`
	InsurerName       string           `json:"insurer_name,omitempty"`
	DepartmentId      string           `json:"department_id,omitempty"`
	SubmissionDate    time.Time        `json:"submission_date,omitempty"`
	ValidUntil        time.Time        `json:"valid_until,omitempty"`
	QuotationNumber   string           `json:"quotation_number,omitempty"`
	PlanName          string           `json:"plan_name,omitempty"`
	CreatedByUserId   string           `json:"created_by_user_id,omitempty"`
	PlanId            string           `json:"plan_id,omitempty"`
}
