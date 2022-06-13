package history

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
)


func (History *History) Delete() {
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("History")

	filter := bson.M{
		"movie_id": bson.M{"$eq": History.MovieID},
		"account_id": bson.M{"$eq": History.AccountID},
	}

	update := bson.M{
		"$set": bson.M{"continue": false,},
	}

	collection.UpdateOne(ctx, filter, update)
}