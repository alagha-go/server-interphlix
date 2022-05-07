package handler

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)


var (
	Client *mongo.Client
	Router = mux.NewRouter()
)

func Main() {

}