package casts

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
)


func Main() {
	LoadCasts()
	Listener()
}


func LoadCasts() {
	var documents []interface{}
	ctx := context.Background()
	collection1 := variables.Client.Database("Interphlix").Collection("Casts")
	collection := variables.Client1.Database("Interphlix").Collection("Casts")

	cursor, err := collection1.Find(ctx, bson.M{})
	variables.HandleError(err, "casts", "LoadCasts", "error while getting casts from the database")
	cursor.All(ctx, &documents)
	collection.Drop(ctx)
	collection.InsertMany(ctx, documents)
}