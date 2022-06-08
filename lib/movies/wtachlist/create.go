package wtachlist

import (
	"context"
	"interphlix/lib/variables"
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