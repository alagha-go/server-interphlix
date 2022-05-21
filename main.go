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
	PORT = ":7000"
)

func main() {
	ConnectToDB()
	go handler.Main()

	err := http.ListenAndServe(PORT, handler.Router)
	HandlError(err)
}

/// connect to both the local and remote mongodb databases
func ConnectToDB() {
	secret := variables.LoadSecret()
	clientOptions := options.Client().
		ApplyURI(secret.RemoteDBUrl)
	ctx := context.Background()
	
	client, err := mongo.Connect(ctx, clientOptions)
	variables.Client = client
	HandlError(err)
}


// handle errors that need the program to exit
func HandlError(err error) {
	if err != nil {
		log.Panic(err)
	}
}