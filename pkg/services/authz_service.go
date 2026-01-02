package services

import (
	"context"
	"github.com/newage-saint/insuretech-go-sdk/pkg/models"
	"strings"
)

// AuthzService handles authz-related API calls
type AuthzService struct {
	Client Client
}

// GetUserRoles GetUserRoles
func (s *AuthzService) GetUserRoles(ctx context.Context, userId string) (*models.UserRolesRetrievalResponse, error) {
	path := "/v1/users/{user_id}/roles"
	path = strings.ReplaceAll(path, "{user_id}", userId)
	var result models.UserRolesRetrievalResponse
	err := s.Client.DoRequest(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// AssignRole AssignRole
func (s *AuthzService) AssignRole(ctx context.Context, userId string, req *models.RoleAssignmentRequest) (*models.RoleAssignmentResponse, error) {
	path := "/v1/users/{user_id}/roles"
	path = strings.ReplaceAll(path, "{user_id}", userId)
	var result models.RoleAssignmentResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateRole CreateRole
func (s *AuthzService) CreateRole(ctx context.Context, req *models.RoleCreationRequest) (*models.RoleCreationResponse, error) {
	path := "/v1/roles"
	var result models.RoleCreationResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CheckPermission CheckPermission
func (s *AuthzService) CheckPermission(ctx context.Context, req *models.CheckPermissionRequest) (*models.CheckPermissionResponse, error) {
	path := "/v1/authz/permissions:check"
	var result models.CheckPermissionResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// EvaluatePolicy EvaluatePolicy
func (s *AuthzService) EvaluatePolicy(ctx context.Context, req *models.PolicyEvaluationRequest) (*models.PolicyEvaluationResponse, error) {
	path := "/v1/authz/policies:evaluate"
	var result models.PolicyEvaluationResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// RevokeRole RevokeRole
func (s *AuthzService) RevokeRole(ctx context.Context, userId string, roleId string) error {
	path := "/v1/users/{user_id}/roles/{role_id}"
	path = strings.ReplaceAll(path, "{user_id}", userId)
	path = strings.ReplaceAll(path, "{role_id}", roleId)
	return s.Client.DoRequest(ctx, "DELETE", path, nil, nil)
}
