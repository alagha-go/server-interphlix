package handler

import (
	"interphlix/lib/handler/movies"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)


var (
	Client *mongo.Client
	Router = mux.NewRouter()
)

func Main() {
	Router.HandleFunc("/", movies.GetHomeMovies)
}