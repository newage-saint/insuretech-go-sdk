package models

// PolicyHealthDeclaration represents a policy_health_declaration
type PolicyHealthDeclaration struct {
	IsSmoker                 bool     `json:"is_smoker,omitempty"`
	BloodGroup               string   `json:"blood_group,omitempty"`
	HasPreExistingConditions bool     `json:"has_pre_existing_conditions,omitempty"`
	Conditions               []string `json:"conditions,omitempty"`
}
