package movies

import (
	"context"
	"interphlix/lib/movies/history"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func GetMyHistory(AccountID primitive.ObjectID, round int) ([]byte, int) {
	var Histories []history.History
	var Movies []Movie
	start := 0
	end := 30
	if round != 0 {
		start = round * 30
		end = round * 30 + 30
	}

	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")

	Histories, err := history.GetLatestHistory(AccountID, false, start, end)
	if err != nil {
		return variables.JsonMarshal(variables.Error{Error: "could not get data"}), http.StatusInternalServerError
	}

	for _, History := range Histories {
		var movie Movie
		err := collection.FindOne(ctx, bson.M{"_id": History.MovieID}).Decode(&movie)
		if err == nil {
			Movies = append(Movies, Movie{ID: movie.ID, Code: movie.Code, Title: movie.Title, Type: movie.Type, ImageUrl: movie.ImageUrl})
		}
	}


	if start >= len(Movies) {
		return []byte(`{"error": "end"}`), http.StatusOK
	}

	if end >= len(Movies) {
		return variables.JsonMarshal(Movies[start:]), http.StatusOK
	}

	return variables.JsonMarshal(Movies[start:end]), http.StatusOK
}


func GetContinue(AccountID primitive.ObjectID, start int, end int) []Movie {
	var Histories []history.History
	var Movies []Movie
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")

	Histories, _ = history.GetLatestHistory(AccountID, true, start, end)

	for index, History := range Histories {
		var movie Movie
		if index < start {
			continue
		}else if index > end {
			return Movies
		}
		err := collection.FindOne(ctx, bson.M{"_id": History.MovieID}).Decode(&movie)
		if err == nil {
			Movies = append(Movies, Movie{ID: movie.ID, Code: movie.Code, Title: movie.Title, Type: movie.Type, ImageUrl: movie.ImageUrl})
		}
	}

	return Movies
}