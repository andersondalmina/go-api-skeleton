package models

// User model
type User struct {
	ID       string `bson:"_id,omitempty" json:"_id,omitempty"`
	Name     string `bson:"name" json:"name"`
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
}
