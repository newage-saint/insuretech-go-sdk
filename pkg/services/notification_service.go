package services

import (
	"context"
	"github.com/newage-saint/insuretech-go-sdk/pkg/models"
	"strings"
)

// NotificationService handles notification-related API calls
type NotificationService struct {
	Client Client
}

// SendNotification SendNotification
func (s *NotificationService) SendNotification(ctx context.Context, req *models.NotificationSendingRequest) (*models.NotificationSendingResponse, error) {
	path := "/v1/notifications"
	var result models.NotificationSendingResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetUserNotifications GetUserNotifications
func (s *NotificationService) GetUserNotifications(ctx context.Context, userId string) (*models.UserNotificationsRetrievalResponse, error) {
	path := "/v1/users/{user_id}/notifications"
	path = strings.ReplaceAll(path, "{user_id}", userId)
	var result models.UserNotificationsRetrievalResponse
	err := s.Client.DoRequest(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// DeactivateNotificationTemplate DeactivateNotificationTemplate
func (s *NotificationService) DeactivateNotificationTemplate(ctx context.Context, templateId string, req *models.NotificationTemplateDeactivationRequest) (*models.NotificationTemplateDeactivationResponse, error) {
	path := "/v1/notification-templates/{template_id}:deactivate"
	path = strings.ReplaceAll(path, "{template_id}", templateId)
	var result models.NotificationTemplateDeactivationResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// SendBulkNotifications SendBulkNotifications
func (s *NotificationService) SendBulkNotifications(ctx context.Context, req *models.BulkNotificationsSendingRequest) (*models.BulkNotificationsSendingResponse, error) {
	path := "/v1/notifications:send-bulk"
	var result models.BulkNotificationsSendingResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// MarkAsRead MarkAsRead
func (s *NotificationService) MarkAsRead(ctx context.Context, req *models.MarkAsReadRequest) (*models.MarkAsReadResponse, error) {
	path := "/v1/notifications:mark-as-read"
	var result models.MarkAsReadResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetNotificationStatus GetNotificationStatus
func (s *NotificationService) GetNotificationStatus(ctx context.Context, notificationId string) (*models.NotificationStatusRetrievalResponse, error) {
	path := "/v1/notifications/{notification_id}"
	path = strings.ReplaceAll(path, "{notification_id}", notificationId)
	var result models.NotificationStatusRetrievalResponse
	err := s.Client.DoRequest(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdatePreferences UpdatePreferences
func (s *NotificationService) UpdatePreferences(ctx context.Context, userId string, req *models.PreferencesUpdateRequest) (*models.PreferencesUpdateResponse, error) {
	path := "/v1/users/{user_id}/notification-preferences"
	path = strings.ReplaceAll(path, "{user_id}", userId)
	var result models.PreferencesUpdateResponse
	err := s.Client.DoRequest(ctx, "PUT", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateNotificationTemplate UpdateNotificationTemplate
func (s *NotificationService) UpdateNotificationTemplate(ctx context.Context, templateId string, req *models.NotificationTemplateUpdateRequest) (*models.NotificationTemplateUpdateResponse, error) {
	path := "/v1/notification-templates/{template_id}"
	path = strings.ReplaceAll(path, "{template_id}", templateId)
	var result models.NotificationTemplateUpdateResponse
	err := s.Client.DoRequest(ctx, "PATCH", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateNotificationTemplate CreateNotificationTemplate
func (s *NotificationService) CreateNotificationTemplate(ctx context.Context, req *models.NotificationTemplateCreationRequest) (*models.NotificationTemplateCreationResponse, error) {
	path := "/v1/notification-templates"
	var result models.NotificationTemplateCreationResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
