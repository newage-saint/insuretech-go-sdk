package services

import (
	"context"
	"github.com/newage-saint/insuretech-go-sdk/pkg/models"
	"strings"
)

// BeneficiaryService handles beneficiary-related API calls
type BeneficiaryService struct {
	Client Client
}

// ListBeneficiaries List beneficiaries (admin)
func (s *BeneficiaryService) ListBeneficiaries(ctx context.Context) (*models.BeneficiariesListingResponse, error) {
	path := "/v1/beneficiaries"
	var result models.BeneficiariesListingResponse
	err := s.Client.DoRequest(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CompleteKYC Complete KYC
func (s *BeneficiaryService) CompleteKYC(ctx context.Context, beneficiaryId string, req *models.KYCCompletionRequest) (*models.KYCCompletionResponse, error) {
	path := "/v1/beneficiaries/{beneficiary_id}/kyc"
	path = strings.ReplaceAll(path, "{beneficiary_id}", beneficiaryId)
	var result models.KYCCompletionResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateIndividualBeneficiary Create individual beneficiary
func (s *BeneficiaryService) CreateIndividualBeneficiary(ctx context.Context, req *models.IndividualBeneficiaryCreationRequest) (*models.IndividualBeneficiaryCreationResponse, error) {
	path := "/v1/beneficiaries/individual"
	var result models.IndividualBeneficiaryCreationResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateBusinessBeneficiary Create business beneficiary
func (s *BeneficiaryService) CreateBusinessBeneficiary(ctx context.Context, req *models.BusinessBeneficiaryCreationRequest) (*models.BusinessBeneficiaryCreationResponse, error) {
	path := "/v1/beneficiaries/business"
	var result models.BusinessBeneficiaryCreationResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetBeneficiary Get beneficiary details
func (s *BeneficiaryService) GetBeneficiary(ctx context.Context, beneficiaryId string) (*models.BeneficiaryRetrievalResponse, error) {
	path := "/v1/beneficiaries/{beneficiary_id}"
	path = strings.ReplaceAll(path, "{beneficiary_id}", beneficiaryId)
	var result models.BeneficiaryRetrievalResponse
	err := s.Client.DoRequest(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateBeneficiary Update beneficiary
func (s *BeneficiaryService) UpdateBeneficiary(ctx context.Context, beneficiaryId string, req *models.BeneficiaryUpdateRequest) (*models.BeneficiaryUpdateResponse, error) {
	path := "/v1/beneficiaries/{beneficiary_id}"
	path = strings.ReplaceAll(path, "{beneficiary_id}", beneficiaryId)
	var result models.BeneficiaryUpdateResponse
	err := s.Client.DoRequest(ctx, "PATCH", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateRiskScore Update risk score
func (s *BeneficiaryService) UpdateRiskScore(ctx context.Context, beneficiaryId string, req *models.RiskScoreUpdateRequest) (*models.RiskScoreUpdateResponse, error) {
	path := "/v1/beneficiaries/{beneficiary_id}/risk-score"
	path = strings.ReplaceAll(path, "{beneficiary_id}", beneficiaryId)
	var result models.RiskScoreUpdateResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
