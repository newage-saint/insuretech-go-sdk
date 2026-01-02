package models

import (
	"time"
)

// UnderwritingHealthDeclaration represents a underwriting_health_declaration
type UnderwritingHealthDeclaration struct {
	HeightCm                 int        `json:"height_cm,omitempty"`
	WeightKg                 string     `json:"weight_kg,omitempty"`
	IsCurrentlyHospitalized  bool       `json:"is_currently_hospitalized,omitempty"`
	MedicalExamDate          time.Time  `json:"medical_exam_date,omitempty"`
	QuoteId                  string     `json:"quote_id"`
	HasFamilyHistory         bool       `json:"has_family_history,omitempty"`
	FamilyHistory            string     `json:"family_history,omitempty"`
	AuditInfo                *AuditInfo `json:"audit_info,omitempty"`
	Id                       string     `json:"id"`
	Bmi                      string     `json:"bmi,omitempty"`
	HasPreExistingConditions bool       `json:"has_pre_existing_conditions,omitempty"`
	PreExistingConditions    string     `json:"pre_existing_conditions,omitempty"`
	Smoker                   bool       `json:"smoker,omitempty"`
	OccupationRiskLevel      string     `json:"occupation_risk_level,omitempty"`
	MedicalDocuments         string     `json:"medical_documents,omitempty"`
	AlcoholConsumer          bool       `json:"alcohol_consumer,omitempty"`
	MedicalExamRequired      bool       `json:"medical_exam_required,omitempty"`
	MedicalExamCompleted     bool       `json:"medical_exam_completed,omitempty"`
	MedicalExamResults       string     `json:"medical_exam_results,omitempty"`
}
