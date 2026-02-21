package services

import (
	"context"
	"github.com/newage-saint/insuretech-go-sdk/pkg/models"
	"strings"
)

// ProductService handles product-related API calls
type ProductService struct {
	Client Client
}

// GetProduct Get product details
func (s *ProductService) GetProduct(ctx context.Context, productId string) (*models.ProductRetrievalResponse, error) {
	path := "/v1/products/{product_id}"
	path = strings.ReplaceAll(path, "{product_id}", productId)
	var result models.ProductRetrievalResponse
	err := s.Client.DoRequest(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateProduct Update product (admin)
func (s *ProductService) UpdateProduct(ctx context.Context, productId string, req *models.ProductUpdateRequest) (*models.ProductUpdateResponse, error) {
	path := "/v1/products/{product_id}"
	path = strings.ReplaceAll(path, "{product_id}", productId)
	var result models.ProductUpdateResponse
	err := s.Client.DoRequest(ctx, "PATCH", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListProducts List all active products
func (s *ProductService) ListProducts(ctx context.Context) (*models.ProductsListingResponse, error) {
	path := "/v1/products"
	var result models.ProductsListingResponse
	err := s.Client.DoRequest(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateProduct Create product (admin)
func (s *ProductService) CreateProduct(ctx context.Context, req *models.ProductCreationRequest) (*models.ProductCreationResponse, error) {
	path := "/v1/products"
	var result models.ProductCreationResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ActivateProduct Activate product
func (s *ProductService) ActivateProduct(ctx context.Context, productId string, req *models.ProductActivationRequest) (*models.ProductActivationResponse, error) {
	path := "/v1/products/{product_id}:activate"
	path = strings.ReplaceAll(path, "{product_id}", productId)
	var result models.ProductActivationResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CalculatePremium Calculate premium
func (s *ProductService) CalculatePremium(ctx context.Context, productId string, req *models.PremiumCalculationRequest) (*models.PremiumCalculationResponse, error) {
	path := "/v1/products/{product_id}:calculate-premium"
	path = strings.ReplaceAll(path, "{product_id}", productId)
	var result models.PremiumCalculationResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeactivateProduct Deactivate product
func (s *ProductService) DeactivateProduct(ctx context.Context, productId string, req *models.ProductDeactivationRequest) (*models.ProductDeactivationResponse, error) {
	path := "/v1/products/{product_id}:deactivate"
	path = strings.ReplaceAll(path, "{product_id}", productId)
	var result models.ProductDeactivationResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DiscontinueProduct Discontinue product
func (s *ProductService) DiscontinueProduct(ctx context.Context, productId string, req *models.DiscontinueProductRequest) (*models.DiscontinueProductResponse, error) {
	path := "/v1/products/{product_id}:discontinue"
	path = strings.ReplaceAll(path, "{product_id}", productId)
	var result models.DiscontinueProductResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// SearchProducts Search products
func (s *ProductService) SearchProducts(ctx context.Context) (*models.ProductsSearchResponse, error) {
	path := "/v1/products:search"
	var result models.ProductsSearchResponse
	err := s.Client.DoRequest(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
