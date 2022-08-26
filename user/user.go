package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type User struct {
	ID             primitive.ObjectID `json:"id" bson:"_id"`
	Username       string             `json:"username" bson:"username"`
	HashedPassword string             `json:"hashed_password" bson:"hashed_password"`
	Role           Role               `json:"role" bson:"role"`
}
