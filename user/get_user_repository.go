package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserByUsername(db *mongo.Database) func(context.Context, string) (*User, error) {
	return func(ctx context.Context, username string) (*User, error) {
		collection := getUsersCollection(db)
		filter := bson.M{"username": username}
		var user User
		if err := collection.FindOne(ctx, filter).Decode(&user); err != nil {
			return nil, err
		}
		return &user, nil
	}
}
