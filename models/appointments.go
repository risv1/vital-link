package models

type Appointments struct {
	ID string `json:"id" bson:"_id"`
	PatientId string `json:"patientId" bson:"patientId"`
	DoctorId string `json:"doctorId" bson:"doctorId"`
	Date string `json:"date" bson:"date"`
	Time string `json:"time" bson:"time"`
	Reason string `json:"reason" bson:"reason"`
	Status string `json:"status" bson:"status"`
	Notes string `json:"notes" bson:"notes"`
}