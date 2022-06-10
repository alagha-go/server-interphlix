package accounts

import (
	"context"
	"interphlix/lib/variables"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


func Main() {
	LoadAccounts()
	Listener()
}

func LoadAccounts() {
	ctx := context.Background()
	var documents []interface{}
	collection1 := variables.Client.Database("Interphlix").Collection("Accounts")
	collection := variables.Client1.Database("Interphlix").Collection("Accounts")

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