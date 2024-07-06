package models

type ContactInformation struct {
	Email string `json:"email" bson:"email"`
	Phone string `json:"phone" bson:"phone"`
	Address Address `json:"address" bson:"address"`
}