package genres

import (
	"context"
	"interphlix/lib/variables"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


func LoadGenres() {
	var documents []interface{}
	ctx := context.Background()
	collection1 := variables.Client.Database("Interphlix").Collection("Genres")
	collection := variables.Client1.Database("Interphlix").Collection("Genres")

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


func (genre *Genre) AddToDB() {
	var Genre Genre
	collection := variables.Client1.Database("Interphlix").Collection("Genres")
	err := collection.FindOne(context.Background(), bson.M{"title": genre.Title}).Decode(&Genre)
	if err != nil {
		collection.InsertOne(context.Background(), genre)
		return
	}
	filter := bson.M{
		"title": genre.Title,
	}
	update := bson.M{"$set": bson.M{
		"types": genre.Types,
	}}
	collection.UpdateOne(context.Background(), filter, update)
}