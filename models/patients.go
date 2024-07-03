package models

type Patient struct {
	ID string `json:"id" bson:"_id"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName string `json:"lastName" bson:"lastName"`
	DateOfBirth string `json:"dateOfBirth" bson:"dateOfBirth"`
	Gender string `json:"gender" bson:"gender"`
	ContactInformation ContactInformation `json:"contactInformation" bson:"contactInformation"`
	MedicalHistory []MedicalHistory `json:"medicalHistory" bson:"medicalHistory"`
	Appointments []string `json:"appointments" bson:"appointments"`
}