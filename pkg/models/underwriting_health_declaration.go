package models

import (
	"time"
)

// UnderwritingHealthDeclaration represents a underwriting_health_declaration
type UnderwritingHealthDeclaration struct {
	Bmi                      string      `json:"bmi,omitempty"`
	MedicalExamRequired      bool        `json:"medical_exam_required,omitempty"`
	MedicalDocuments         string      `json:"medical_documents,omitempty"`
	Id                       string      `json:"id"`
	QuoteId                  string      `json:"quote_id"`
	IsCurrentlyHospitalized  bool        `json:"is_currently_hospitalized,omitempty"`
	HasFamilyHistory         bool        `json:"has_family_history,omitempty"`
	FamilyHistory            string      `json:"family_history,omitempty"`
	Smoker                   bool        `json:"smoker,omitempty"`
	AlcoholConsumer          bool        `json:"alcohol_consumer,omitempty"`
	WeightKg                 string      `json:"weight_kg,omitempty"`
	HasPreExistingConditions bool        `json:"has_pre_existing_conditions,omitempty"`
	OccupationRiskLevel      string      `json:"occupation_risk_level,omitempty"`
	MedicalExamResults       string      `json:"medical_exam_results,omitempty"`
	MedicalExamDate          time.Time   `json:"medical_exam_date,omitempty"`
	AuditInfo                interface{} `json:"audit_info"`
	PreExistingConditions    string      `json:"pre_existing_conditions,omitempty"`
	MedicalExamCompleted     bool        `json:"medical_exam_completed,omitempty"`
	HeightCm                 int         `json:"height_cm,omitempty"`
}
