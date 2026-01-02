package services

import (
	"context"
	"github.com/newage-saint/insuretech-go-sdk/pkg/models"
	"strings"
)

// PolicyService handles policy-related API calls
type PolicyService struct {
	Client Client
}

// GetPolicy GetPolicy
func (s *PolicyService) GetPolicy(ctx context.Context, policyId string) (*models.PolicyRetrievalResponse, error) {
	path := "/v1/policies/{policy_id}"
	path = strings.ReplaceAll(path, "{policy_id}", policyId)
	var result models.PolicyRetrievalResponse
	err := s.Client.DoRequest(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdatePolicy UpdatePolicy
func (s *PolicyService) UpdatePolicy(ctx context.Context, policyId string, req *models.PolicyUpdateRequest) (*models.PolicyUpdateResponse, error) {
	path := "/v1/policies/{policy_id}"
	path = strings.ReplaceAll(path, "{policy_id}", policyId)
	var result models.PolicyUpdateResponse
	err := s.Client.DoRequest(ctx, "PATCH", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CancelPolicy CancelPolicy
func (s *PolicyService) CancelPolicy(ctx context.Context, policyId string, req *models.PolicyCancellationRequest) (*models.PolicyCancellationResponse, error) {
	path := "/v1/policies/{policy_id}:cancel"
	path = strings.ReplaceAll(path, "{policy_id}", policyId)
	var result models.PolicyCancellationResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// RenewPolicy RenewPolicy
func (s *PolicyService) RenewPolicy(ctx context.Context, policyId string, req *models.PolicyPolicyRenewalRequest) (*models.PolicyPolicyRenewalResponse, error) {
	path := "/v1/policies/{policy_id}:renew"
	path = strings.ReplaceAll(path, "{policy_id}", policyId)
	var result models.PolicyPolicyRenewalResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GeneratePolicyDocument GeneratePolicyDocument
func (s *PolicyService) GeneratePolicyDocument(ctx context.Context, policyId string, req *models.PolicyDocumentGenerationRequest) (*models.PolicyDocumentGenerationResponse, error) {
	path := "/v1/policies/{policy_id}:generate-document"
	path = strings.ReplaceAll(path, "{policy_id}", policyId)
	var result models.PolicyDocumentGenerationResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreatePolicy CreatePolicy
func (s *PolicyService) CreatePolicy(ctx context.Context, req *models.PolicyCreationRequest) (*models.PolicyCreationResponse, error) {
	path := "/v1/policies"
	var result models.PolicyCreationResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// IssuePolicy IssuePolicy
func (s *PolicyService) IssuePolicy(ctx context.Context, policyId string, req *models.PolicyIssuanceRequest) (*models.PolicyIssuanceResponse, error) {
	path := "/v1/policies/{policy_id}:issue"
	path = strings.ReplaceAll(path, "{policy_id}", policyId)
	var result models.PolicyIssuanceResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListUserPolicies ListUserPolicies
func (s *PolicyService) ListUserPolicies(ctx context.Context, customerId string) (*models.UserPoliciesListingResponse, error) {
	path := "/v1/users/{customer_id}/policies"
	path = strings.ReplaceAll(path, "{customer_id}", customerId)
	var result models.UserPoliciesListingResponse
	err := s.Client.DoRequest(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
