package models 

type MedicalHistory struct {
	Condition string `json:"condition" bson:"condition"`
	DiagnosisDate string `json:"diagnosisDate" bson:"diagnosisDate"`
	Notes string `json:"notes" bson:"notes"`
}