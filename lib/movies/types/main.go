package types

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
)

func Main() {
	LoadTypes()
	Listener()
}


func LoadTypes() {
	var documents []interface{}
	ctx := context.Background()
	collection1 := variables.Client.Database("Interphlix").Collection("Types")
	collection := variables.Client1.Database("Interphlix").Collection("Types")

	cursor, err := collection1.Find(ctx, bson.M{})
	variables.HandleError(err, "types", "LoadTypes", "error while loading data from the database")
	err = cursor.All(ctx, &documents)
	variables.HandleError(err, "types", "LoadTypes", "error decoding cursor")
	collection.Drop(ctx)
	_, err = collection.InsertMany(ctx, documents)
	variables.HandleError(err, "types", "LoadTypes", "error inserting types to the local database")
}


func TypeExists(Type string) bool {
	var DbType interface{}
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Types")
	err := collection.FindOne(ctx, bson.M{"type": Type}).Decode(&DbType)
	return err == nil
}