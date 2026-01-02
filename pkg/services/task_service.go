package services

import (
	"context"
	"github.com/newage-saint/insuretech-go-sdk/pkg/models"
	"strings"
)

// TaskService handles task-related API calls
type TaskService struct {
	Client Client
}

// ListMyTasks ListMyTasks
func (s *TaskService) ListMyTasks(ctx context.Context) (*models.MyTasksListingResponse, error) {
	path := "/v1/tasks/my-tasks"
	var result models.MyTasksListingResponse
	err := s.Client.DoRequest(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateTask CreateTask
func (s *TaskService) CreateTask(ctx context.Context, req *models.TaskCreationRequest) (*models.TaskCreationResponse, error) {
	path := "/v1/tasks"
	var result models.TaskCreationResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetTask GetTask
func (s *TaskService) GetTask(ctx context.Context, taskId string) (*models.TaskRetrievalResponse, error) {
	path := "/v1/tasks/{task_id}"
	path = strings.ReplaceAll(path, "{task_id}", taskId)
	var result models.TaskRetrievalResponse
	err := s.Client.DoRequest(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateTask UpdateTask
func (s *TaskService) UpdateTask(ctx context.Context, taskId string, req *models.TaskUpdateRequest) (*models.TaskUpdateResponse, error) {
	path := "/v1/tasks/{task_id}"
	path = strings.ReplaceAll(path, "{task_id}", taskId)
	var result models.TaskUpdateResponse
	err := s.Client.DoRequest(ctx, "PATCH", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// AssignTask AssignTask
func (s *TaskService) AssignTask(ctx context.Context, taskId string, req *models.TaskAssignmentRequest) (*models.TaskAssignmentResponse, error) {
	path := "/v1/tasks/{task_id}"
	path = strings.ReplaceAll(path, "{task_id}", taskId)
	var result models.TaskAssignmentResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
