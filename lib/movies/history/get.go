package history

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetMovieHistory(AccountID, MovieID primitive.ObjectID) string {
	var History History
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("History")

	err := collection.FindOne(ctx, bson.M{"account_id": AccountID, "movie_id": MovieID}).Decode(&History)
	if err != nil {
		return `{"error": "no history for the provided movie"}`
	}
	return string(variables.JsonMarshal(History))
}