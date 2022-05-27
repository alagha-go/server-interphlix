package ratings

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
)


func Main() {
	LoadRatings()
	Listener()
}


func LoadRatings() {
	var documents []interface{}
	ctx := context.Background()
	collection1 := variables.Client.Database("Interphlix").Collection("Ratings")
	collection := variables.Client1.Database("Interphlix").Collection("Ratings")

	cursor, err := collection1.Find(ctx, bson.M{})
	if err != nil {
		variables.HandleError(err, "ratings", "LoadRatings", "error while getting ratings from the remote database")
		return
	}
	cursor.All(ctx, &documents)
	collection.Drop(ctx)
	_, err = collection.InsertMany(ctx, documents)
	variables.HandleError(err, "ratings", "LoadRatings", "error while inserting ratings to the local database")
}