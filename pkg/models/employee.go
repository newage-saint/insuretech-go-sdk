package models

import (
	"time"
)

// Employee represents a employee
type Employee struct {
	CoverageAmount    *Money          `json:"coverage_amount,omitempty"`
	PremiumAmount     *Money          `json:"premium_amount,omitempty"`
	Status            *EmployeeStatus `json:"status,omitempty"`
	CreatedAt         time.Time       `json:"created_at,omitempty"`
	UpdatedAt         time.Time       `json:"updated_at,omitempty"`
	EmployeeUuid      string          `json:"employee_uuid,omitempty"`
	Name              string          `json:"name,omitempty"`
	EmployeeId        string          `json:"employee_id,omitempty"`
	DepartmentId      string          `json:"department_id,omitempty"`
	BusinessId        string          `json:"business_id,omitempty"`
	InsuranceCategory *InsuranceType  `json:"insurance_category,omitempty"`
	AssignedPlanId    string          `json:"assigned_plan_id,omitempty"`
}
