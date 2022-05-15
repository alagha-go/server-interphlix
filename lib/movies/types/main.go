package types

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
)


var (
	Types []Type
)


func Main() {

}


func LoadTypes() {
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Types")

	cursor, err := collection.Find(ctx, bson.M{})
	variables.HandleError(err, "types", "LoadTypes", "error while loading data from the database")
	cursor.All(ctx, &Types)
}