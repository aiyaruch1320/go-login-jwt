package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
	CreatedAt      time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at" bson:"updated_at"`
}

type UserReq struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Role     Role   `json:"role" bson:"role"`
}

func (userReq *UserReq) mapToUser() *User {
	return &User{
		Username:       userReq.Username,
		HashedPassword: HashedPassword(userReq.Password),
		Role:           userReq.Role,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}
