package models

// InsurerRetrievalResponse represents a insurer_retrieval_response
type InsurerRetrievalResponse struct {
	Config  *InsurerConfig `json:"config,omitempty"`
	Error   *Error         `json:"error,omitempty"`
	Insurer *Insurer       `json:"insurer,omitempty"`
}
