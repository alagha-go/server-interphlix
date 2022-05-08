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
	ClientOptions := options.Client().ApplyURI(secret.RemoteDBUrl)
	ctx := context.Background()
	
	variables.Client, err = mongo.Connect(ctx, ClientOptions)
	HandlError(err)
}


// handle errors that need the program to exit
func HandlError(err error) {
	if err != nil {
		log.Panic(err)
	}
}