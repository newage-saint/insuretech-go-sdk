package services

import (
	"context"
	"github.com/newage-saint/insuretech-go-sdk/pkg/models"
	"strings"
)

// DocumentService handles document-related API calls
type DocumentService struct {
	Client Client
}

// ListDocuments ListDocuments
func (s *DocumentService) ListDocuments(ctx context.Context, entityType string, entityId string) (*models.DocumentsListingResponse, error) {
	path := "/v1/entities/{entity_type}/{entity_id}/documents"
	path = strings.ReplaceAll(path, "{entity_type}", entityType)
	path = strings.ReplaceAll(path, "{entity_id}", entityId)
	var result models.DocumentsListingResponse
	err := s.Client.DoRequest(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeactivateDocumentTemplate DeactivateDocumentTemplate
func (s *DocumentService) DeactivateDocumentTemplate(ctx context.Context, templateId string, req *models.DocumentTemplateDeactivationRequest) (*models.DocumentTemplateDeactivationResponse, error) {
	path := "/v1/document-templates/{template_id}:deactivate"
	path = strings.ReplaceAll(path, "{template_id}", templateId)
	var result models.DocumentTemplateDeactivationResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetDocumentTemplate GetDocumentTemplate
func (s *DocumentService) GetDocumentTemplate(ctx context.Context, templateId string) (*models.DocumentTemplateRetrievalResponse, error) {
	path := "/v1/document-templates/{template_id}"
	path = strings.ReplaceAll(path, "{template_id}", templateId)
	var result models.DocumentTemplateRetrievalResponse
	err := s.Client.DoRequest(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateDocumentTemplate UpdateDocumentTemplate
func (s *DocumentService) UpdateDocumentTemplate(ctx context.Context, templateId string, req *models.DocumentTemplateUpdateRequest) (*models.DocumentTemplateUpdateResponse, error) {
	path := "/v1/document-templates/{template_id}"
	path = strings.ReplaceAll(path, "{template_id}", templateId)
	var result models.DocumentTemplateUpdateResponse
	err := s.Client.DoRequest(ctx, "PATCH", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateDocumentTemplate CreateDocumentTemplate
func (s *DocumentService) CreateDocumentTemplate(ctx context.Context, req *models.DocumentTemplateCreationRequest) (*models.DocumentTemplateCreationResponse, error) {
	path := "/v1/document-templates"
	var result models.DocumentTemplateCreationResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListDocumentTemplates ListDocumentTemplates
func (s *DocumentService) ListDocumentTemplates(ctx context.Context) (*models.DocumentTemplatesListingResponse, error) {
	path := "/v1/document-templates"
	var result models.DocumentTemplatesListingResponse
	err := s.Client.DoRequest(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DownloadDocument DownloadDocument
func (s *DocumentService) DownloadDocument(ctx context.Context, documentId string) (*models.DocumentDownloadResponse, error) {
	path := "/v1/documents/{document_id}/download"
	path = strings.ReplaceAll(path, "{document_id}", documentId)
	var result models.DocumentDownloadResponse
	err := s.Client.DoRequest(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetDocument GetDocument
func (s *DocumentService) GetDocument(ctx context.Context, documentId string) (*models.DocumentRetrievalResponse, error) {
	path := "/v1/documents/{document_id}"
	path = strings.ReplaceAll(path, "{document_id}", documentId)
	var result models.DocumentRetrievalResponse
	err := s.Client.DoRequest(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GenerateDocument GenerateDocument
func (s *DocumentService) GenerateDocument(ctx context.Context, req *models.DocumentGenerationRequest) (*models.DocumentGenerationResponse, error) {
	path := "/v1/documents"
	var result models.DocumentGenerationResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
