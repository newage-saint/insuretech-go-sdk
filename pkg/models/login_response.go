package models

// LoginResponse represents a login_response
type LoginResponse struct {
	Error                 *Error `json:"error,omitempty"`
	UserId                string `json:"user_id,omitempty"`
	SessionId             string `json:"session_id,omitempty"`
	SessionToken          string `json:"session_token,omitempty"`
	AccessToken           string `json:"access_token,omitempty"`
	RefreshToken          string `json:"refresh_token,omitempty"`
	AccessTokenExpiresIn  int    `json:"access_token_expires_in,omitempty"`
	RefreshTokenExpiresIn int    `json:"refresh_token_expires_in,omitempty"`
	CsrfToken             string `json:"csrf_token,omitempty"`
	User                  *User  `json:"user,omitempty"`
	SessionType           string `json:"session_type,omitempty"`
}
