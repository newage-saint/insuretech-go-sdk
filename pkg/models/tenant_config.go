package models

// TenantConfig represents a tenant_config
type TenantConfig struct {
	Id          string      `json:"id"`
	TenantId    string      `json:"tenant_id"`
	ConfigKey   string      `json:"config_key"`
	ConfigValue string      `json:"config_value"`
	Type        *ConfigType `json:"type"`
	Description string      `json:"description,omitempty"`
	AuditInfo   *AuditInfo  `json:"audit_info,omitempty"`
}
