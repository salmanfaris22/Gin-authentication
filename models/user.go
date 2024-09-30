package models

type User struct {
	ID        string `bson:"_id,omitempty" json:"id"` // Use string for MongoDB ObjectID
	FirstName string `bson:"first_name" json:"first_name"`
	LastName  string `bson:"last_name" json:"last_name"`
	Email     string `bson:"email" json:"email"`
	Password  string `bson:"password" json:"password"`
	Token     string `json:"token"`
}
