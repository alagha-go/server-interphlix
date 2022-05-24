package casts

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
)


var (
	Casts []Cast
)

func Main() {
	LoadCasts()
	Listener()
}


func LoadCasts() {
	collection := variables.Client.Database("Interphlix").Collection("Casts")

	cursor, err := collection.Find(context.Background(), bson.M{})
	variables.HandleError(err, "casts", "LoadCasts", "error while getting casts from the database")
	cursor.All(context.Background(), &Casts)
}