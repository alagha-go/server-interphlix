package history

import (
	"context"
	"interphlix/lib/variables"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func (History *History) Create() string {
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("History")

	History.LastTimeWatched = time.Now()
	History.Episode.LastTimeWatched = time.Now()
	History.Episodes = append(History.Episodes, *History.Episode)
	History.Episode = nil
	History.ID = primitive.NewObjectID()

	_, err := collection.InsertOne(ctx, History)
	if err != nil {
		variables.HandleError(err, "history", "History.Create", "error while inserting document to the database")
		return `{"error": "could not add movie to watchlist"}`
	}
	return "success"
}


func (history *History) Exists() bool {
	var History History
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("History")

	err := collection.FindOne(ctx, bson.M{"account_id": history.AccountID, "movie_id": history.MovieID}).Decode(&History)
	return err == nil
}