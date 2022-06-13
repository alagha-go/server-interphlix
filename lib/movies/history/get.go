package history

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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


func GetLatestHistory(AccountID primitive.ObjectID, Continue bool, start, end int) ([]History, error) {
	var Histories []History
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("History")

	sort := bson.D{{"last_time_watched", -1}}
	opts := options.Find().SetSort(sort)

	if Continue {
		cursor, err := collection.Find(ctx, bson.M{"account_id": AccountID, "continue": true}, opts)
		if err != nil {
			variables.HandleError(err, "history", "GetLatestHistory", "error while getting data from the database")
		}
		cursor.All(ctx, &Histories)
	}else {
		cursor, err := collection.Find(ctx, bson.M{"account_id": AccountID}, opts)
		if err != nil {
			variables.HandleError(err, "history", "GetLatestHistory", "error while getting data from the database")
		}
		cursor.All(ctx, &Histories)
	}

	if len(Histories) < start {
		return []History{}, nil
	}else if len(Histories) < end {
		return Histories[start:], nil
	}

	return Histories[start:end], nil
}