package services

import (
	"context"
	"github.com/newage-saint/insuretech-go-sdk/pkg/models"
	"strings"
)

// EndorsementService handles endorsement-related API calls
type EndorsementService struct {
	Client Client
}

// RequestEndorsement RequestEndorsement
func (s *EndorsementService) RequestEndorsement(ctx context.Context, policyId string, req *models.RequestEndorsementRequest) (*models.RequestEndorsementResponse, error) {
	path := "/v1/policies/{policy_id}/endorsements"
	path = strings.ReplaceAll(path, "{policy_id}", policyId)
	var result models.RequestEndorsementResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListEndorsements ListEndorsements
func (s *EndorsementService) ListEndorsements(ctx context.Context, policyId string) (*models.EndorsementsListingResponse, error) {
	path := "/v1/policies/{policy_id}/endorsements"
	path = strings.ReplaceAll(path, "{policy_id}", policyId)
	var result models.EndorsementsListingResponse
	err := s.Client.DoRequest(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetEndorsement GetEndorsement
func (s *EndorsementService) GetEndorsement(ctx context.Context, endorsementId string) (*models.EndorsementRetrievalResponse, error) {
	path := "/v1/endorsements/{endorsement_id}"
	path = strings.ReplaceAll(path, "{endorsement_id}", endorsementId)
	var result models.EndorsementRetrievalResponse
	err := s.Client.DoRequest(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ApproveEndorsement ApproveEndorsement
func (s *EndorsementService) ApproveEndorsement(ctx context.Context, endorsementId string, req *models.EndorsementApprovalRequest) (*models.EndorsementApprovalResponse, error) {
	path := "/v1/endorsements/{endorsement_id}"
	path = strings.ReplaceAll(path, "{endorsement_id}", endorsementId)
	var result models.EndorsementApprovalResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
