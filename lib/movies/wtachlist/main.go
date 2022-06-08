package wtachlist

import (
	"context"
	"interphlix/lib/variables"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


func Main() {

}


func LoadWatchList() {
	var documents []interface{}
	ctx := context.Background()
	collection1 := variables.Client.Database("Interphlix").Collection("Watchlist")
	collection := variables.Client1.Database("Interphlix").Collection("Watchlist")

	cursor, err := collection1.Find(ctx, bson.M{})
	if err != nil {
		log.Panic(err)
	}
	err = cursor.All(ctx, &documents)
	if err != nil {
		log.Panic(err)
	}

	_, err = collection.InsertMany(ctx, documents)
	if err != nil && err != mongo.ErrEmptySlice {
		log.Panic(err)
	}
}