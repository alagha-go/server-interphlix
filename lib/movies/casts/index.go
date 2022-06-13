package casts

import (
	"context"
	"interphlix/lib/variables"
	"log"

	"github.com/blevesearch/bleve/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


var (
	Index bleve.Index
)


func StartIndex() {
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Casts")

	model := mongo.IndexModel{
		Keys: bson.D{{"title", "text"}},
	}

	_, err := collection.Indexes().CreateOne(ctx, model)
	if err != nil {
		log.Println("Casts")
		log.Panic(err)
	}
}