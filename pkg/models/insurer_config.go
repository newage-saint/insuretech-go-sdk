package models

// InsurerConfig represents a insurer_config
type InsurerConfig struct {
	AuthCredentials           string              `json:"auth_credentials,omitempty"`
	WebhookSecret             string              `json:"webhook_secret,omitempty"`
	BusinessModel             string              `json:"business_model,omitempty"`
	AuditInfo                 interface{}         `json:"audit_info"`
	ApiVersion                string              `json:"api_version,omitempty"`
	AuthType                  *AuthenticationType `json:"auth_type,omitempty"`
	RealTimeClaimNotification bool                `json:"real_time_claim_notification,omitempty"`
	WebhookUrl                string              `json:"webhook_url,omitempty"`
	UnderwritingThreshold     int                 `json:"underwriting_threshold,omitempty"`
	ClaimSettlementDays       int                 `json:"claim_settlement_days,omitempty"`
	AutoUnderwritingEnabled   bool                `json:"auto_underwriting_enabled,omitempty"`
	PaymentTerms              string              `json:"payment_terms,omitempty"`
	Id                        string              `json:"id"`
	InsurerId                 string              `json:"insurer_id"`
	ApiBaseUrl                string              `json:"api_base_url,omitempty"`
}
