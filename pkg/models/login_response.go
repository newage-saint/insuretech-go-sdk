package models

// LoginResponse represents a login_response
type LoginResponse struct {
	SessionId             string `json:"session_id,omitempty"`
	AccessToken           string `json:"access_token,omitempty"`
	AccessTokenExpiresIn  int    `json:"access_token_expires_in,omitempty"`
	SessionToken          string `json:"session_token,omitempty"`
	UserId                string `json:"user_id,omitempty"`
	RefreshToken          string `json:"refresh_token,omitempty"`
	RefreshTokenExpiresIn int    `json:"refresh_token_expires_in,omitempty"`
	CsrfToken             string `json:"csrf_token,omitempty"`
	User                  *User  `json:"user,omitempty"`
	SessionType           string `json:"session_type,omitempty"`
	Error                 *Error `json:"error,omitempty"`
}
