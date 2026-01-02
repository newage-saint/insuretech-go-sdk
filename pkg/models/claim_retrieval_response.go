package models

// ClaimRetrievalResponse represents a claim_retrieval_response
type ClaimRetrievalResponse struct {
	Error *Error `json:"error,omitempty"`
	Claim *Claim `json:"claim,omitempty"`
}
