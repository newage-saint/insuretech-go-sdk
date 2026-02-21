package models

// TransactionRetrievalResponse represents a transaction_retrieval_response
type TransactionRetrievalResponse struct {
	Error       *Error          `json:"error,omitempty"`
	Transaction *MFSTransaction `json:"transaction,omitempty"`
}
