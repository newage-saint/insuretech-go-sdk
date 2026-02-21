package models

// PolicyRetrievalResponse represents a policy_retrieval_response
type PolicyRetrievalResponse struct {
	Error  *Error  `json:"error,omitempty"`
	Policy *Policy `json:"policy,omitempty"`
}
