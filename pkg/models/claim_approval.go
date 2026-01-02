package models

import (
	"time"
)

// ClaimApproval represents a claim_approval
type ClaimApproval struct {
	ClaimId        string      `json:"claim_id"`
	ApproverId     string      `json:"approver_id"`
	Decision       interface{} `json:"decision"`
	ApprovedAmount *Money      `json:"approved_amount,omitempty"`
	Notes          string      `json:"notes,omitempty"`
	ApprovalId     string      `json:"approval_id"`
	ApproverRole   string      `json:"approver_role"`
	ApprovalLevel  int         `json:"approval_level"`
	ApprovedAt     time.Time   `json:"approved_at,omitempty"`
	CreatedAt      time.Time   `json:"created_at"`
}
