package history

import (
	"context"
	"interphlix/lib/variables"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func (History *History) Create() {
	if History.Exists() {
		History.Update()
		return
	}
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("History")

	History.LastTimeWatched = time.Now()
	History.Episode.LastTimeWatched = time.Now()
	History.Episodes = append(History.Episodes, *History.Episode)
	History.LastSeasonIndex = History.Episode.SeasonIndex
	History.Continue = true
	History.Episode = nil
	History.ID = primitive.NewObjectID()

	_, err := collection.InsertOne(ctx, History)
	if err != nil {
		variables.HandleError(err, "history", "History.Create", "error while inserting document to the database")
		return
	}
}


func (History *History) LocalCreate() {
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("History")

	collection.InsertOne(ctx, History)
}




func (history *History) Exists() bool {
	var History History
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("History")

	err := collection.FindOne(ctx, bson.M{"account_id": history.AccountID, "movie_id": history.MovieID}).Decode(&History)
	history.ID = History.ID
	return err == nil
}