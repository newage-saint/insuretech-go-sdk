package models

// PartnerRetrievalResponse represents a partner_retrieval_response
type PartnerRetrievalResponse struct {
	Error   *Error   `json:"error,omitempty"`
	Partner *Partner `json:"partner,omitempty"`
}
