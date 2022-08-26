package config

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DB struct {
	Client *mongo.Client
}

func InitMongoDB(ctx context.Context) *DB {
	// To configure auth via URI instead of a Credential, use
	// "mongodb://user:password@localhost:27017".
	credential := options.Credential{
		AuthMechanism: "SCRAM-SHA-1",
		AuthSource:    "test",
		Username:      "root",
		Password:      "password",
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credential))
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	conn := &DB{
		Client: client,
	}

	return conn
}
