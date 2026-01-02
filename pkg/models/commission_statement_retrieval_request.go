package models

// CommissionStatementRetrievalRequest represents a commission_statement_retrieval_request
type CommissionStatementRetrievalRequest struct {
	PeriodStart string `json:"period_start,omitempty"`
	PeriodEnd   string `json:"period_end,omitempty"`
	RecipientId string `json:"recipient_id"`
}
