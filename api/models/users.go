package models 

type User struct {
	ID string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Role string `json:"role" bson:"role"`
	ProfileId string `json:"profileId" bson:"profileId"`
}