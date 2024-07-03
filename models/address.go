package models

type Address struct {
	Street string `json:"street" bson:"street"`
	City string `json:"city" bson:"city"`
	State string `json:"state" bson:"state"`
	Pincode string `json:"pincode" bson:"pincode"`
	Country string `json:"country" bson:"country"`
}