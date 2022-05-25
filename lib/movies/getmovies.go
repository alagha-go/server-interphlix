package movies

import (
	"context"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func GetMoviesByGenre(genre string) ([]byte, int) {
	var Movies []Movie
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Movies")

	cursor, err := collection.Find(ctx, bson.M{"genres": bson.A{genre}})
	variables.HandleError(err, "movies", "GetMovieByGenre", "error while getting movies from the database")
	err = cursor.All(ctx, &Movies)
	variables.HandleError(err, "movies", "GetMovieByGenre", "error while decoding cursor")
	return variables.JsonMarshal(Movies), http.StatusOK
}