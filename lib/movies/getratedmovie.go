package movies

import (
	"context"
	"interphlix/lib/movies/ratings"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)



func GetRatedMovies(AccountID primitive.ObjectID, round int) ([]byte, int) {
	var Ratings []ratings.Rate
	var RatedMovies []RatedMovie
	start := 0
	end := 30
	ctx := context.Background()
	collection1 := variables.Client1.Database("Interphlix").Collection("Ratings")
	collection := variables.Client1.Database("Interphlix").Collection("Movies")

	cursor, err := collection1.Find(ctx, bson.M{"account_id": AccountID})
	if err != nil {
		variables.HandleError(err, "movies", "GetRatedMovies", "error while getting ratings from the database")
		return variables.JsonMarshal(variables.Error{Error: "could not get ratings from the database"}), http.StatusInternalServerError
	}
	err = cursor.All(ctx, &Ratings)
	if err != nil {
		variables.HandleError(err, "movies", "GetRatedMovies", "error while decoding ratings from the cursor")
		return variables.JsonMarshal(variables.Error{Error: "could not decode data"}), http.StatusInternalServerError
	}
	
	for index := range Ratings {
		var RatedMovie RatedMovie
		RatedMovie.Rate = Ratings[index]
		collection.FindOne(ctx, bson.M{"_id": Ratings[index].MovieID}).Decode(&RatedMovie.Movie)
		RatedMovie.Movie = Movie{ID: RatedMovie.Movie.ID, Code: RatedMovie.Movie.Code, Title: RatedMovie.Movie.Title, Type: RatedMovie.Movie.Type, ImageUrl: RatedMovie.Movie.ImageUrl}
		RatedMovies = append(RatedMovies, RatedMovie)
	}

	if round != 0 {
		start = round * 30
		end = round * 30 + 30
	}

	if start >= len(RatedMovies) {
		return []byte(`{"error": "end"}`), http.StatusOK
	}

	if end >= len(RatedMovies) {
		return variables.JsonMarshal(RatedMovies[start:]), http.StatusOK
	}

	return variables.JsonMarshal(RatedMovies[start:end]), http.StatusOK
}