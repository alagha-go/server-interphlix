package watchlist

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
)


func (WatchList *WatchList) Delete() string {
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Watchlist")

	_, err := collection.DeleteOne(ctx, bson.M{"movie_id": WatchList.MovieID, "account_id": WatchList.AccountID})
	if err != nil {
		variables.HandleError(err, "watchlist", "Watchlist.Delete", "could not delete watchlist")
		return `{"error": "could not delete watch list"}`
	}
	return `{"deleted": true}`
}