package models

// ProductCreationResponse represents a product_creation_response
type ProductCreationResponse struct {
	Message   string `json:"message,omitempty"`
	Error     *Error `json:"error,omitempty"`
	ProductId string `json:"product_id,omitempty"`
}
