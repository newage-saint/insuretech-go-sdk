package models

import (
	"time"
)

// UnderwritingHealthDeclaration represents a underwriting_health_declaration
type UnderwritingHealthDeclaration struct {
	PreExistingConditions    string      `json:"pre_existing_conditions,omitempty"`
	OccupationRiskLevel      string      `json:"occupation_risk_level,omitempty"`
	MedicalExamResults       string      `json:"medical_exam_results,omitempty"`
	MedicalDocuments         string      `json:"medical_documents,omitempty"`
	QuoteId                  string      `json:"quote_id"`
	HeightCm                 int         `json:"height_cm,omitempty"`
	Smoker                   bool        `json:"smoker,omitempty"`
	HasPreExistingConditions bool        `json:"has_pre_existing_conditions,omitempty"`
	IsCurrentlyHospitalized  bool        `json:"is_currently_hospitalized,omitempty"`
	HasFamilyHistory         bool        `json:"has_family_history,omitempty"`
	AlcoholConsumer          bool        `json:"alcohol_consumer,omitempty"`
	MedicalExamRequired      bool        `json:"medical_exam_required,omitempty"`
	MedicalExamCompleted     bool        `json:"medical_exam_completed,omitempty"`
	AuditInfo                interface{} `json:"audit_info"`
	Bmi                      string      `json:"bmi,omitempty"`
	FamilyHistory            string      `json:"family_history,omitempty"`
	MedicalExamDate          time.Time   `json:"medical_exam_date,omitempty"`
	Id                       string      `json:"id"`
	WeightKg                 string      `json:"weight_kg,omitempty"`
}
