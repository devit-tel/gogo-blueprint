package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewConnection(clientOptions *options.ClientOptions) (*mongo.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	db, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return db, nil
}
