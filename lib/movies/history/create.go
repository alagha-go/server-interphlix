package history

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
)


func (history *History) Exists() bool {
	var History History
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("History")

	err := collection.FindOne(ctx, bson.M{"account_id": history.AccountID, "movie_id": history.MovieID}).Decode(&History)
	return err == nil
}