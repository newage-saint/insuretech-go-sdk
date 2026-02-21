package models

import (
	"time"
)

// Department represents a department
type Department struct {
	DepartmentId string    `json:"department_id,omitempty"`
	Name         string    `json:"name,omitempty"`
	BusinessId   string    `json:"business_id,omitempty"`
	EmployeeNo   int       `json:"employee_no,omitempty"`
	TotalPremium *Money    `json:"total_premium,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}
