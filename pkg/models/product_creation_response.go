package models

// ProductCreationResponse represents a product_creation_response
type ProductCreationResponse struct {
	Error     *Error `json:"error,omitempty"`
	ProductId string `json:"product_id,omitempty"`
	Message   string `json:"message,omitempty"`
}
