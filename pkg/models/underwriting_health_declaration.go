package models

import (
	"time"
)

// UnderwritingHealthDeclaration represents a underwriting_health_declaration
type UnderwritingHealthDeclaration struct {
	HasPreExistingConditions bool        `json:"has_pre_existing_conditions,omitempty"`
	MedicalExamCompleted     bool        `json:"medical_exam_completed,omitempty"`
	MedicalExamResults       string      `json:"medical_exam_results,omitempty"`
	MedicalExamDate          time.Time   `json:"medical_exam_date,omitempty"`
	Id                       string      `json:"id"`
	HeightCm                 int         `json:"height_cm,omitempty"`
	WeightKg                 string      `json:"weight_kg,omitempty"`
	PreExistingConditions    string      `json:"pre_existing_conditions,omitempty"`
	FamilyHistory            string      `json:"family_history,omitempty"`
	Smoker                   bool        `json:"smoker,omitempty"`
	MedicalExamRequired      bool        `json:"medical_exam_required,omitempty"`
	QuoteId                  string      `json:"quote_id"`
	HasFamilyHistory         bool        `json:"has_family_history,omitempty"`
	AlcoholConsumer          bool        `json:"alcohol_consumer,omitempty"`
	OccupationRiskLevel      string      `json:"occupation_risk_level,omitempty"`
	MedicalDocuments         string      `json:"medical_documents,omitempty"`
	Bmi                      string      `json:"bmi,omitempty"`
	IsCurrentlyHospitalized  bool        `json:"is_currently_hospitalized,omitempty"`
	AuditInfo                interface{} `json:"audit_info"`
}
