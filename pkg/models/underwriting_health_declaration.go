package models

import (
	"time"
)

// UnderwritingHealthDeclaration represents a underwriting_health_declaration
type UnderwritingHealthDeclaration struct {
	PreExistingConditions    string      `json:"pre_existing_conditions,omitempty"`
	FamilyHistory            string      `json:"family_history,omitempty"`
	MedicalExamDate          time.Time   `json:"medical_exam_date,omitempty"`
	MedicalDocuments         string      `json:"medical_documents,omitempty"`
	QuoteId                  string      `json:"quote_id"`
	WeightKg                 string      `json:"weight_kg,omitempty"`
	HasPreExistingConditions bool        `json:"has_pre_existing_conditions,omitempty"`
	OccupationRiskLevel      string      `json:"occupation_risk_level,omitempty"`
	HeightCm                 int         `json:"height_cm,omitempty"`
	HasFamilyHistory         bool        `json:"has_family_history,omitempty"`
	AlcoholConsumer          bool        `json:"alcohol_consumer,omitempty"`
	MedicalExamCompleted     bool        `json:"medical_exam_completed,omitempty"`
	MedicalExamResults       string      `json:"medical_exam_results,omitempty"`
	AuditInfo                interface{} `json:"audit_info"`
	Bmi                      string      `json:"bmi,omitempty"`
	IsCurrentlyHospitalized  bool        `json:"is_currently_hospitalized,omitempty"`
	Smoker                   bool        `json:"smoker,omitempty"`
	MedicalExamRequired      bool        `json:"medical_exam_required,omitempty"`
	Id                       string      `json:"id"`
}
