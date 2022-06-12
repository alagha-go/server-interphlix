package casts

import (
	"context"
	"interphlix/lib/variables"
	"log"

	"github.com/blevesearch/bleve/v2"
	"go.mongodb.org/mongo-driver/bson"
)


var (
	Index bleve.Index
	err error
)


func StartIndex() {
	var Casts []Cast
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Casts")

	mapping := bleve.NewIndexMapping()
	Index, err = bleve.New("Casts", mapping)
	if err != nil {
		Index, err = bleve.Open("Casts")
		if err != nil {
			log.Panic(err)
		}
	}

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Panic(err)
	}
	err = cursor.All(ctx, &Casts)
	if err != nil {
		log.Panic(err)
	}

	for _, Cast := range Casts {
		Index.Index(Cast.ID.Hex(), Cast)
	}
}


func (Cast *Cast) AddIndex() {
	Index.Index(Cast.ID.Hex(), Cast)
}