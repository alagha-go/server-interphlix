package movies

import (
	"context"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)


func GetMovies() ([]byte, int) {
	var Movies []Movie
	collection := variables.Client1.Database("Interphlix").Collection("Movies")
	cursor, err := collection.Find(context.Background(), bson.M{})
	variables.HandleError(err, "movies", "GetMovies", "error while getting movies from the local database")
	cursor.All(context.Background(), &Movies)
	return variables.JsonMarshal(Movies), http.StatusOK
}