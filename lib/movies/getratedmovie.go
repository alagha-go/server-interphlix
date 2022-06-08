package movies

import (
	"context"
	"interphlix/lib/movies/ratings"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)



func GetRatedMovies(AccountID primitive.ObjectID) ([]byte, int) {
	var Ratings []ratings.Rate
	var RatedMovies []RatedMovie
	ctx := context.Background()
	collection1 := variables.Client1.Database("Interphlix").Collection("Ratings")
	collection := variables.Client1.Database("Interphlix").Collection("Movies")

	cursor, err := collection1.Find(ctx, bson.M{"account_id": AccountID})
	if err != nil {
		variables.HandleError(err, "movies", "GetRatedMovies", "error while getting ratings from the database")
		return variables.JsonMarshal(variables.Error{Error: "could not get ratings from the database"}), http.StatusInternalServerError
	}
	err = cursor.All(ctx, Ratings)
	if err != nil {
		variables.HandleError(err, "movies", "GetRatedMovies", "error while decoding ratings from the cursor")
		return variables.JsonMarshal(variables.Error{Error: "could not decode data"}), http.StatusInternalServerError
	}
	
	for index := range Ratings {
		var RatedMovie RatedMovie
		RatedMovie.Rate = Ratings[index]
		collection.FindOne(ctx, bson.M{"_id": Ratings[index].MovieID}).Decode(&RatedMovie.Movie)
		RatedMovies = append(RatedMovies, RatedMovie)
	}

	return variables.JsonMarshal(RatedMovies), http.StatusOK
}