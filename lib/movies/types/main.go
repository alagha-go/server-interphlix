package types

import (
	"context"
	"interphlix/lib/variables"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
	if err != nil {
		log.Panic(err)
	}
	err = cursor.All(ctx, &documents)
	if err != nil {
		log.Panic(err)
	}

	collection.Drop(ctx)
	_, err = collection.InsertMany(ctx, documents)
	if err != nil && err != mongo.ErrEmptySlice {
		log.Panic(err)
	}
}


func TypeExists(Type string) bool {
	var DbType interface{}
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Types")
	err := collection.FindOne(ctx, bson.M{"type": Type}).Decode(&DbType)
	return err == nil
}