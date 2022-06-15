package movies

import (
	"context"
	"interphlix/lib/movies/watchlist"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func GetMyWatchlist(AccountID primitive.ObjectID, round int) ([]byte, int) {
	var WatchLists []watchlist.WatchList
	var Movies []Movie
	start := 0
	end := 30
	ctx := context.Background()
	collection1 := variables.Client1.Database("Interphlix").Collection("Watchlist")
	collection := variables.Client1.Database("Interphlix").Collection("Movies")

	cursor, err := collection1.Find(ctx, bson.M{"account_id": AccountID})
	if err != nil {
		variables.HandleError(err, "movies", "GetMyWatchlist", "error while getting watchlists from the database")
		return variables.JsonMarshal(variables.Error{Error: "could not get watchlist"}), http.StatusInternalServerError
	}
	err = cursor.All(ctx, &WatchLists)
	if err != nil {
		variables.HandleError(err, "movies", "GetMyWatchlist", "error while decoding cursor")
		return variables.JsonMarshal(variables.Error{Error: "could not get watchlist"}), http.StatusInternalServerError
	}

	for _, watchlist := range WatchLists {
		var movie Movie
		err := collection.FindOne(ctx, bson.M{"_id": watchlist.MovieID}).Decode(&movie)
		if err == nil {
			Movies = append(Movies, Movie{ID: movie.ID, Code: movie.Code, Title: movie.Title, Type: movie.Type, ImageUrl: movie.ImageUrl})
		}
	}

	if round != 0 {
		start = round * 30
		end = round * 30 + 30
	}

	if start >= len(Movies) {
		return []byte(`{"error": "end"}`), http.StatusOK
	}

	if len(Movies) >= end {
		return variables.JsonMarshal(Movies[start:]), http.StatusOK
	}

	return variables.JsonMarshal(Movies[start:end]), http.StatusOK
}