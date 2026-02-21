package models

// ProductRetrievalResponse represents a product_retrieval_response
type ProductRetrievalResponse struct {
	Error   *Error   `json:"error,omitempty"`
	Product *Product `json:"product,omitempty"`
}
