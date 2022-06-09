package history

import (
	"context"
	"interphlix/lib/variables"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)


func (history *History) Update() {
	var History History
	ctx := context.Background()
	collection1 := variables.Client.Database("Interphlix").Collection("History")
	collection := variables.Client1.Database("Interphlix").Collection("History")

	err := collection.FindOne(ctx, bson.M{"account_id": history.AccountID, "movie_id": history.MovieID}).Decode(&History)
	if err != nil {
		history.Create()
		return
	}

	if history.Episode != nil {
		for index := range History.Episodes {
			if History.Episodes[index].ID == history.Episode.ID {
				History.Episodes[index].Percentage = history.Episode.Percentage
				History.Episodes[index].LastTimeWatched = time.Now()
			}
		}
	}else {
		History.Percentage = history.Percentage
	}

	filter := bson.M{"_id": bson.M{"$eq": history.ID}}

	update := bson.M{"$set": bson.M{
		"episodes": History.Episodes,
		"last_time_watched": time.Now(),
		"percentage": History.Percentage,
	}}

	_, err = collection1.UpdateOne(ctx, filter, update)
	if err != nil {
		variables.HandleError(err, "history", "history.Update", "error while updating document")
		return
	}
}


func (History *History) LocalUpdate() {
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("History")

	filter := bson.M{"_id": bson.M{"$eq": History.ID}}
	update := bson.M{"$set": History}

	collection.UpdateOne(ctx, filter, update)
}