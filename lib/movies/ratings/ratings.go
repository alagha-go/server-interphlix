package ratings

import (
	"context"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func GetMovieRatings(MovieID primitive.ObjectID) ([]byte, int) {
	var Ratings []Rate
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Ratings")

	cursor, err := collection.Find(ctx, bson.M{"movie_id": MovieID})
	if err != nil {
		variables.HandleError(err, "ratings", "GetMovieRatings", "error while getting data from the database")
		return variables.JsonMarshal(variables.Error{Error: "could not get ratings"}), http.StatusInternalServerError
	}
	err = cursor.All(ctx, &Ratings)
	if err != nil {
		variables.HandleError(err, "ratings", "GetMovieRatings", "error while decoding cursor")
		return variables.JsonMarshal(variables.Error{Error: "could not decode data"}), http.StatusInternalServerError
	}
	return variables.JsonMarshal(Ratings), http.StatusOK
}