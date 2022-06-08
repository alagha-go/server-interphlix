package wtachlist

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
)


func (WatchList *WatchList) Create() string {
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Watchlist")

	_, err := collection.InsertOne(ctx, WatchList)
	if err != nil {
		variables.HandleError(err, "watchlist", "WatchList.Create", "error while inserting document to the database")
		return string(variables.JsonMarshal(variables.Error{Error: "could not create watchlist"}))
	}
	return `watchlist created`
}


func (WatchList *WatchList) Exists() bool {
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Watchlist")

	err := collection.FindOne(ctx, bson.M{"account_id": WatchList.AccountID, "movie_id": WatchList.MovieID}).Decode(WatchList)
	return err == nil
}