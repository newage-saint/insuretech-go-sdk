package services

import (
	"context"
	"github.com/newage-saint/insuretech-go-sdk/pkg/models"
	"strings"
)

// AuthService handles auth-related API calls
type AuthService struct {
	Client Client
}

// Register Register new user
func (s *AuthService) Register(ctx context.Context, req *models.RegistrationRequest) (*models.RegistrationResponse, error) {
	path := "/v1/auth/register"
	var result models.RegistrationResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// RefreshToken Refresh access token
func (s *AuthService) RefreshToken(ctx context.Context, req *models.RefreshTokenRequest) (*models.RefreshTokenResponse, error) {
	path := "/v1/auth/token:refresh"
	var result models.RefreshTokenResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// VerifyOTP Verify OTP
func (s *AuthService) VerifyOTP(ctx context.Context, req *models.OTPVerificationRequest) (*models.OTPVerificationResponse, error) {
	path := "/v1/auth/otp:verify"
	var result models.OTPVerificationResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// RevokeAllSessions Revoke all sessions for a user (logout from all devices)
func (s *AuthService) RevokeAllSessions(ctx context.Context, userId string, req *models.RevokeAllSessionsRequest) (*models.RevokeAllSessionsResponse, error) {
	path := "/v1/auth/users/{user_id}/sessions:revoke-all"
	path = strings.ReplaceAll(path, "{user_id}", userId)
	var result models.RevokeAllSessionsResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// SendOTP Send OTP for verification
func (s *AuthService) SendOTP(ctx context.Context, req *models.OTPSendingRequest) (*models.OTPSendingResponse, error) {
	path := "/v1/auth/otp:send"
	var result models.OTPSendingResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Login Login with credentials
func (s *AuthService) Login(ctx context.Context, req *models.LoginRequest) (*models.LoginResponse, error) {
	path := "/v1/auth/login"
	var result models.LoginResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetSession Get session details
func (s *AuthService) GetSession(ctx context.Context, sessionId string) (*models.SessionRetrievalResponse, error) {
	path := "/v1/auth/sessions/{session_id}"
	path = strings.ReplaceAll(path, "{session_id}", sessionId)
	var result models.SessionRetrievalResponse
	err := s.Client.DoRequest(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// RevokeSession Revoke a specific session
func (s *AuthService) RevokeSession(ctx context.Context, sessionId string) error {
	path := "/v1/auth/sessions/{session_id}"
	path = strings.ReplaceAll(path, "{session_id}", sessionId)
	return s.Client.DoRequest(ctx, "DELETE", path, nil, nil)
}

// ValidateCSRF Validate CSRF token (server-side sessions only)
func (s *AuthService) ValidateCSRF(ctx context.Context, req *models.CSRFValidationRequest) (*models.CSRFValidationResponse, error) {
	path := "/v1/auth/csrf:validate"
	var result models.CSRFValidationResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ResetPassword Reset password
func (s *AuthService) ResetPassword(ctx context.Context, req *models.ResetPasswordRequest) (*models.ResetPasswordResponse, error) {
	path := "/v1/auth/password:reset"
	var result models.ResetPasswordResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// GetCurrentSession Get current user's active session
func (s *AuthService) GetCurrentSession(ctx context.Context) (*models.CurrentSessionRetrievalResponse, error) {
	path := "/v1/auth/session/current"
	var result models.CurrentSessionRetrievalResponse
	err := s.Client.DoRequest(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ListSessions List all sessions for a user
func (s *AuthService) ListSessions(ctx context.Context, userId string) (*models.SessionsListingResponse, error) {
	path := "/v1/auth/users/{user_id}/sessions"
	path = strings.ReplaceAll(path, "{user_id}", userId)
	var result models.SessionsListingResponse
	err := s.Client.DoRequest(ctx, "GET", path, nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ChangePassword Change password
func (s *AuthService) ChangePassword(ctx context.Context, req *models.ChangePasswordRequest) (*models.ChangePasswordResponse, error) {
	path := "/v1/auth/password:change"
	var result models.ChangePasswordResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// ValidateToken Validate token
func (s *AuthService) ValidateToken(ctx context.Context, req *models.TokenValidationRequest) (*models.TokenValidationResponse, error) {
	path := "/v1/auth/token:validate"
	var result models.TokenValidationResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Logout Logout
func (s *AuthService) Logout(ctx context.Context, req *models.LogoutRequest) (*models.LogoutResponse, error) {
	path := "/v1/auth/logout"
	var result models.LogoutResponse
	err := s.Client.DoRequest(ctx, "POST", path, req, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
