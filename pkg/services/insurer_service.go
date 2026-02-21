package services

import (
	"context"
	"github.com/newage-saint/insuretech-go-sdk/pkg/models"
	"strings"
)

// InsurerService handles insurer-related API calls
type InsurerService struct {
	Client Client
}

// UpdateInsurerConfig Update insurer config
func (s *InsurerService) UpdateInsurerConfig(ctx context.Context, insurerId string, req *models.InsurerConfigUpdateRequest) (*models.InsurerConfigUpdateResponse, error) {
	path := "/v1/insurers/{insurer_id}/config"
	path = strings.ReplaceAll(path, "{insurer_id}", insurerId)
	var result models.InsurerConfigUpdateResponse
	err := s.Client.DoRequest(ctx, "PUT", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// AddInsurerProduct Add insurer product
func (s *InsurerService) AddInsurerProduct(ctx context.Context, insurerId string, req *models.AddInsurerProductRequest) (*models.AddInsurerProductResponse, error) {
	path := "/v1/insurers/{insurer_id}/products"
	path = strings.ReplaceAll(path, "{insurer_id}", insurerId)
	var result models.AddInsurerProductResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListInsurerProducts List insurer products
func (s *InsurerService) ListInsurerProducts(ctx context.Context, insurerId string) (*models.InsurerProductsListingResponse, error) {
	path := "/v1/insurers/{insurer_id}/products"
	path = strings.ReplaceAll(path, "{insurer_id}", insurerId)
	var result models.InsurerProductsListingResponse
	err := s.Client.DoRequest(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetInsurer Get insurer details
func (s *InsurerService) GetInsurer(ctx context.Context, insurerId string) (*models.InsurerRetrievalResponse, error) {
	path := "/v1/insurers/{insurer_id}"
	path = strings.ReplaceAll(path, "{insurer_id}", insurerId)
	var result models.InsurerRetrievalResponse
	err := s.Client.DoRequest(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateInsurer Update insurer
func (s *InsurerService) UpdateInsurer(ctx context.Context, insurerId string, req *models.InsurerUpdateRequest) (*models.InsurerUpdateResponse, error) {
	path := "/v1/insurers/{insurer_id}"
	path = strings.ReplaceAll(path, "{insurer_id}", insurerId)
	var result models.InsurerUpdateResponse
	err := s.Client.DoRequest(ctx, "PATCH", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetInsurerProduct Get insurer product
func (s *InsurerService) GetInsurerProduct(ctx context.Context, insurerProductId string) (*models.InsurerProductRetrievalResponse, error) {
	path := "/v1/insurer-products/{insurer_product_id}"
	path = strings.ReplaceAll(path, "{insurer_product_id}", insurerProductId)
	var result models.InsurerProductRetrievalResponse
	err := s.Client.DoRequest(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateInsurerProduct Update insurer product
func (s *InsurerService) UpdateInsurerProduct(ctx context.Context, insurerProductId string, req *models.InsurerProductUpdateRequest) (*models.InsurerProductUpdateResponse, error) {
	path := "/v1/insurer-products/{insurer_product_id}"
	path = strings.ReplaceAll(path, "{insurer_product_id}", insurerProductId)
	var result models.InsurerProductUpdateResponse
	err := s.Client.DoRequest(ctx, "PATCH", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateInsurer Create insurer
func (s *InsurerService) CreateInsurer(ctx context.Context, req *models.InsurerCreationRequest) (*models.InsurerCreationResponse, error) {
	path := "/v1/insurers"
	var result models.InsurerCreationResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListInsurers List insurers
func (s *InsurerService) ListInsurers(ctx context.Context) (*models.InsurersListingResponse, error) {
	path := "/v1/insurers"
	var result models.InsurersListingResponse
	err := s.Client.DoRequest(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
