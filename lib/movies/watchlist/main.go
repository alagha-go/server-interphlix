package watchlist

import (
	"context"
	"interphlix/lib/variables"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


func Main() {
	LoadWatchList()
	Listener()
	go ListenLength()
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

func ListenLength() {
	ctx := context.Background()
	collection1 := variables.Client.Database("Interphlix").Collection("Watchlist")
	collection := variables.Client1.Database("Interphlix").Collection("Watchlist")

	for {
		count1, err := collection1.CountDocuments(ctx, bson.M{})
		if err != nil {
			variables.HandleError(err, "watchlist", "ListenLength", "error while getting document count")
		}
		count, err := collection.CountDocuments(ctx, bson.M{})
		if err != nil {
			variables.HandleError(err, "watchlist", "ListenLength", "error while getting document count")
		}
		if count1 != count {
			LoadWatchList()
		}
		time.Sleep(time.Second)
	}

}