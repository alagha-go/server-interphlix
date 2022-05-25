package main

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func ConnectToRemoteDB1() {
	secret := variables.LoadSecret()
	clientOptions := options.Client().
		ApplyURI(secret.RemoteDBUrl)
	ctx := context.Background()
	
	client, err := mongo.Connect(ctx, clientOptions)
	variables.Client = client
	HandlError(err)
}