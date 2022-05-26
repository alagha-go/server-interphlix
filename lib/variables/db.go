package variables

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func ConnectToRemoteDB1() {
	secret := LoadSecret()
	clientOptions := options.Client().
		ApplyURI(secret.RemoteDBUrl)
	ctx := context.Background()
	
	client, err := mongo.Connect(ctx, clientOptions)
	Client = client
	handleerror(err)
}


func ConnectRemotedDB2() {
	secret := LoadSecret()
	clientOptions := options.Client().
		ApplyURI(secret.Remote2DBUrl)
	ctx := context.Background()
	
	client, err := mongo.Connect(ctx, clientOptions)
	Client2 = client
	handleerror(err)
}

func ConnectLocalDB() {
	secret := LoadSecret()
	clientOptions := options.Client().
		ApplyURI(secret.LocalDBUrl)
	ctx := context.Background()
	
	client, err := mongo.Connect(ctx, clientOptions)
	Client1 = client
	handleerror(err)
}

func handleerror(err error) {
	if err != nil {
		log.Panic(err)
	}
}