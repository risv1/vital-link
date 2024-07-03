package models 

type Records struct {
	ID string `json:"id" bson:"_id"`
	PatientId string `json:"patientId" bson:"patientId"`
	DoctorId string `json:"doctorId" bson:"doctorId"`
	Date string `json:"date" bson:"date"`
	Time string `json:"time" bson:"time"`
	Type string `json:"type" bson:"type"`
	Description string `json:"description" bson:"description"`
	Documents []Documents `json:"documents" bson:"documents"`
}

type Documents struct {
	DocumentId string `json:"documentId" bson:"documentId"`
	DocumentType string `json:"documentType" bson:"documentType"`
	Url string `json:"url" bson:"url"`
}