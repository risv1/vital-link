package models

type Doctor struct {
	ID string `json:"id" bson:"_id"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName string `json:"lastName" bson:"lastName"`
	Specialty string `json:"specialty" bson:"specialty"`
	ContactInformation ContactInformation `json:"contactInformation" bson:"contactInformation"`
	Availability []Availability `json:"availability" bson:"availability"`
	Image string `json:"image" bson:"image"`
}

type Availability struct {
	Day string `json:"day" bson:"day"`
	StartTime string `json:"startTime" bson:"startTime"`
	EndTime string `json:"endTime" bson:"endTime"`
}
