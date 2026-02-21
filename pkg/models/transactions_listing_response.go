package models

// TransactionsListingResponse represents a transactions_listing_response
type TransactionsListingResponse struct {
	TotalCount   int               `json:"total_count,omitempty"`
	Error        *Error            `json:"error,omitempty"`
	Transactions []*MFSTransaction `json:"transactions,omitempty"`
}
