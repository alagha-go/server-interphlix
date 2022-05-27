package ratings

import (
	"context"
	"interphlix/lib/movies"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


func RateMovie(Rate Rate) ([]byte, int) {
	Movie := movies.Movie{ID: Rate.MovieID}
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Ratings")
	if Rate.Exists() {
		return variables.JsonMarshal(variables.Error{Error: "rate already exists"}), http.StatusBadRequest
	}
	Rate.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(ctx, Rate)
	if err != nil {
		variables.HandleError(err, "ratings", "RateMovie", "error while inserting rate to the database")
		return variables.JsonMarshal(variables.Error{Error: "could save your rate"}), http.StatusInternalServerError
	}
	Movie.UpdateRate(Rate.Stars)
	return variables.JsonMarshal(Rate), http.StatusOK
}