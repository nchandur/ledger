package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectDB() error {
	uri := "mongodb://localhost:27017"

	ctx := context.TODO()

	var err error

	Client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		return err
	}

	return nil
}

func DisconnectDB() error {
	if err := Client.Disconnect(context.TODO()); err != nil {
		return err
	}
	return nil
}
