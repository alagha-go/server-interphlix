package main

import (
	"context"
	"interphlix/lib/handler"
	"interphlix/lib/variables"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	PORT = ":8000"
)

func main() {
	ConnectToDB()
	go handler.Main()

	err := http.ListenAndServe(PORT, handler.Router)
	HandlError(err)
}

/// connect to both the local and remote mongodb databases
func ConnectToDB() {
	var err error
	secret := variables.LoadSecret()
	LocalClientOptions := options.Client().ApplyURI(secret.LocalDBUrl)
	RemoteClientOptions := options.Client().ApplyURI(secret.RemoteDBUrl)
	ctx := context.Background()
	
	variables.LocalClient, err = mongo.Connect(ctx, LocalClientOptions)
	HandlError(err)
	variables.RemoteClient, err = mongo.Connect(ctx, RemoteClientOptions)
	HandlError(err)
}


// handle errors that need the program to exit
func HandlError(err error) {
	if err != nil {
		log.Panic(err)
	}
}