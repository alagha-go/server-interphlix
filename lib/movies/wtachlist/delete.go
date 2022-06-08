package wtachlist

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
)


func (WatchList *WatchList) Delete() string {
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Watchlist")

	_, err := collection.DeleteOne(ctx, bson.M{"_id": WatchList.ID})
	if err != nil {
		variables.HandleError(err, "watchlist", "Watchlist.Delete", "could not delete watchlist")
		return `{"error": "could not delete watch list"}`
	}
	return `{"deleted": true}`
}