package models

import (
	"time"
)

// ClaimApproval represents a claim_approval
type ClaimApproval struct {
	CreatedAt        time.Time         `json:"created_at,omitempty"`
	ApprovalId       string            `json:"approval_id,omitempty"`
	ApproverId       string            `json:"approver_id,omitempty"`
	ApproverRole     string            `json:"approver_role,omitempty"`
	ApprovalLevel    int               `json:"approval_level,omitempty"`
	Decision         *ApprovalDecision `json:"decision,omitempty"`
	ApprovedAmount   *Money            `json:"approved_amount,omitempty"`
	Notes            string            `json:"notes,omitempty"`
	ApprovedCurrency string            `json:"approved_currency,omitempty"`
	ClaimId          string            `json:"claim_id,omitempty"`
	ApprovedAt       time.Time         `json:"approved_at,omitempty"`
}
