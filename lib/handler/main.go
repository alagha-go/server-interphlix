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
	movies.Main()
	Router.HandleFunc("/", movies.GetHomeMovies)
	Router.HandleFunc("/movies/upload", movies.UploadMovie)
	Router.HandleFunc("/home", movies.GetHomeMovies)
	Router.HandleFunc("/movies/addserver/{id}", movies.UploadMovie)
	Router.HandleFunc("/movies/addurl/{id}/{url}", movies.UploadMovie)
	Router.HandleFunc("/movies/setserver/{id}/{servername}/{serverid}", movies.UploadMovie)
}