package movies

import (
	"context"
	"interphlix/lib/movies/history"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func GetMyHistory(AccountID primitive.ObjectID, round int) ([]byte, int) {
	var Histories []history.History
	var Movies []Movie
	start := 0
	end := 30

	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")
	collection1 := variables.Client1.Database("Interphlix").Collection("History")

	opts := options.Find().SetSort(bson.D{{"last_time_watched", -1}})

	cursor, err := collection1.Find(ctx, bson.M{"account_id": AccountID}, opts)
	if err != nil {
		variables.HandleError(err, "movies", "GetMyHistory", "error while getting sorted history")
		return variables.JsonMarshal(variables.Error{Error: "could not get history"}), http.StatusInternalServerError
	}
	err = cursor.All(ctx, &Histories)
	if err != nil {
		variables.HandleError(err, "movies", "GetMyHistory", "error while decoding cursor")
		return variables.JsonMarshal(variables.Error{Error: "could not get history"}), http.StatusInternalServerError
	}

	for _, History := range Histories {
		var Movie Movie
		err := collection.FindOne(ctx, bson.M{"_id": History.MovieID}).Decode(&Movie)
		if err == nil {
			Movies = append(Movies, Movie)
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

	return variables.JsonMarshal(Movies), http.StatusOK
}