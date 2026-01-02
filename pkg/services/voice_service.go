package services

import (
	"context"
	"github.com/newage-saint/insuretech-go-sdk/pkg/models"
	"strings"
)

// VoiceService handles voice-related API calls
type VoiceService struct {
	Client Client
}

// GetVoiceSession GetVoiceSession
func (s *VoiceService) GetVoiceSession(ctx context.Context, voiceSessionId string) (*models.VoiceSessionRetrievalResponse, error) {
	path := "/v1/voice-sessions/{voice_session_id}"
	path = strings.ReplaceAll(path, "{voice_session_id}", voiceSessionId)
	var result models.VoiceSessionRetrievalResponse
	err := s.Client.DoRequest(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// EndVoiceSession EndVoiceSession
func (s *VoiceService) EndVoiceSession(ctx context.Context, voiceSessionId string, req *models.EndVoiceSessionRequest) (*models.EndVoiceSessionResponse, error) {
	path := "/v1/voice-sessions/{voice_session_id}"
	path = strings.ReplaceAll(path, "{voice_session_id}", voiceSessionId)
	var result models.EndVoiceSessionResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// StartVoiceSession StartVoiceSession
func (s *VoiceService) StartVoiceSession(ctx context.Context, req *models.VoiceSessionStartRequest) (*models.VoiceSessionStartResponse, error) {
	path := "/v1/voice-sessions"
	var result models.VoiceSessionStartResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetTranscript GetTranscript
func (s *VoiceService) GetTranscript(ctx context.Context, voiceSessionId string) (*models.TranscriptRetrievalResponse, error) {
	path := "/v1/voice-sessions/{voice_session_id}/transcript"
	path = strings.ReplaceAll(path, "{voice_session_id}", voiceSessionId)
	var result models.TranscriptRetrievalResponse
	err := s.Client.DoRequest(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ProcessVoiceCommand ProcessVoiceCommand
func (s *VoiceService) ProcessVoiceCommand(ctx context.Context, voiceSessionId string, req *models.VoiceCommandProcessingRequest) (*models.VoiceCommandProcessingResponse, error) {
	path := "/v1/voice-sessions/{voice_session_id}/commands"
	path = strings.ReplaceAll(path, "{voice_session_id}", voiceSessionId)
	var result models.VoiceCommandProcessingResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
