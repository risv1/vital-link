package models 

type User struct {
	GoogleID string `json:"googleId" bson:"googleId"`
	Name string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Role string `json:"role" bson:"role"`
	ProfileId string `json:"profileId" bson:"profileId"`
}