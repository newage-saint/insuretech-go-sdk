package models

// InsurerProductUpdateResponse represents a insurer_product_update_response
type InsurerProductUpdateResponse struct {
	Error   *Error `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}
