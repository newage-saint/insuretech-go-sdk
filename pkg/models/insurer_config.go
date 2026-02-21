package models

// InsurerConfig represents a insurer_config
type InsurerConfig struct {
	WebhookUrl                string              `json:"webhook_url,omitempty"`
	PaymentTerms              string              `json:"payment_terms,omitempty"`
	AuditInfo                 interface{}         `json:"audit_info"`
	BusinessModel             string              `json:"business_model,omitempty"`
	AutoUnderwritingEnabled   bool                `json:"auto_underwriting_enabled,omitempty"`
	ApiBaseUrl                string              `json:"api_base_url,omitempty"`
	WebhookSecret             string              `json:"webhook_secret,omitempty"`
	UnderwritingThreshold     int                 `json:"underwriting_threshold,omitempty"`
	InsurerId                 string              `json:"insurer_id"`
	ApiVersion                string              `json:"api_version,omitempty"`
	AuthCredentials           string              `json:"auth_credentials,omitempty"`
	RealTimeClaimNotification bool                `json:"real_time_claim_notification,omitempty"`
	ClaimSettlementDays       int                 `json:"claim_settlement_days,omitempty"`
	Id                        string              `json:"id"`
	AuthType                  *AuthenticationType `json:"auth_type,omitempty"`
}
