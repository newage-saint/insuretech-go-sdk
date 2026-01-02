package services

import (
	"context"
	"github.com/newage-saint/insuretech-go-sdk/pkg/models"
	"strings"
)

// UnderwritingService handles underwriting-related API calls
type UnderwritingService struct {
	Client Client
}

// GetQuote GetQuote
func (s *UnderwritingService) GetQuote(ctx context.Context, quoteId string) (*models.QuoteRetrievalResponse, error) {
	path := "/v1/quotes/{quote_id}"
	path = strings.ReplaceAll(path, "{quote_id}", quoteId)
	var result models.QuoteRetrievalResponse
	err := s.Client.DoRequest(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ApproveUnderwriting ApproveUnderwriting
func (s *UnderwritingService) ApproveUnderwriting(ctx context.Context, quoteId string, req *models.UnderwritingApprovalRequest) (*models.UnderwritingApprovalResponse, error) {
	path := "/v1/quotes/{quote_id}"
	path = strings.ReplaceAll(path, "{quote_id}", quoteId)
	var result models.UnderwritingApprovalResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetHealthDeclaration GetHealthDeclaration
func (s *UnderwritingService) GetHealthDeclaration(ctx context.Context, quoteId string) (*models.HealthDeclarationRetrievalResponse, error) {
	path := "/v1/quotes/{quote_id}/health-declaration"
	path = strings.ReplaceAll(path, "{quote_id}", quoteId)
	var result models.HealthDeclarationRetrievalResponse
	err := s.Client.DoRequest(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// SubmitHealthDeclaration SubmitHealthDeclaration
func (s *UnderwritingService) SubmitHealthDeclaration(ctx context.Context, quoteId string, req *models.HealthDeclarationSubmissionRequest) (*models.HealthDeclarationSubmissionResponse, error) {
	path := "/v1/quotes/{quote_id}/health-declaration"
	path = strings.ReplaceAll(path, "{quote_id}", quoteId)
	var result models.HealthDeclarationSubmissionResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListQuotes ListQuotes
func (s *UnderwritingService) ListQuotes(ctx context.Context, beneficiaryId string) (*models.QuotesListingResponse, error) {
	path := "/v1/beneficiaries/{beneficiary_id}/quotes"
	path = strings.ReplaceAll(path, "{beneficiary_id}", beneficiaryId)
	var result models.QuotesListingResponse
	err := s.Client.DoRequest(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetUnderwritingDecision GetUnderwritingDecision
func (s *UnderwritingService) GetUnderwritingDecision(ctx context.Context, quoteId string) (*models.UnderwritingDecisionRetrievalResponse, error) {
	path := "/v1/quotes/{quote_id}/decision"
	path = strings.ReplaceAll(path, "{quote_id}", quoteId)
	var result models.UnderwritingDecisionRetrievalResponse
	err := s.Client.DoRequest(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// RequestQuote RequestQuote
func (s *UnderwritingService) RequestQuote(ctx context.Context, req *models.RequestQuoteRequest) (*models.RequestQuoteResponse, error) {
	path := "/v1/quotes"
	var result models.RequestQuoteResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
