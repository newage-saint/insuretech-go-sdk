package models

// ProductDeactivationRequest represents a product_deactivation_request
type ProductDeactivationRequest struct {
	Reason    string `json:"reason,omitempty"`
	ProductId string `json:"product_id"`
}
