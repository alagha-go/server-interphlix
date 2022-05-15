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
	LoadTypes()
}


func LoadTypes() {
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Types")

	cursor, err := collection.Find(ctx, bson.M{})
	variables.HandleError(err, "types", "LoadTypes", "error while loading data from the database")
	cursor.All(ctx, &Types)
}


func TypeExists(Type string) bool {
	for index := range Types {
		if Types[index].Type == Type {
			return true
		}
	}
	return false
}