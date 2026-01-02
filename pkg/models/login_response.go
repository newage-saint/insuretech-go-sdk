package models

// LoginResponse represents a login_response
type LoginResponse struct {
	UserId                string `json:"user_id,omitempty"`
	AccessToken           string `json:"access_token,omitempty"`
	RefreshToken          string `json:"refresh_token,omitempty"`
	AccessTokenExpiresIn  int    `json:"access_token_expires_in,omitempty"`
	RefreshTokenExpiresIn int    `json:"refresh_token_expires_in,omitempty"`
	User                  *User  `json:"user,omitempty"`
	Error                 *Error `json:"error,omitempty"`
}
