package movies

import (
	"context"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)


func GetMoviesByCast() ([]byte, int) {
	var Movies []Movie
	var movies []Movie
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Movies")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		variables.HandleError(err, "movies", "GetMoviesByCast", "error while getting data from the local database")
		return variables.JsonMarshal(variables.Error{Error: "could not get data"}), http.StatusInternalServerError
	}
	cursor.All(ctx, &Movies)

	for _, Movie := range Movies {
		if Movie.ContainsCast() {
			movies = append(movies, Movie)
		}
	}

	return variables.JsonMarshal(movies), http.StatusOK
}