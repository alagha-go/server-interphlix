package movies

import (
	"context"
	"interphlix/lib/variables"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func StartIndex() {
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Movies")

	model := mongo.IndexModel{
		Keys: bson.D{{"title", "text"}, {"casts", "text"}, {"genre", "text"}},
	}

	_, err := collection.Indexes().CreateOne(ctx, model)
	if err != nil {
		log.Println("Movies")
		log.Panic(err)
	}
}