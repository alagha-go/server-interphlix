package movies

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
	var Movies []Movie
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Movies")

	mapping := bleve.NewIndexMapping()
	Index, err = bleve.New("Movies", mapping)
	if err != nil {
		Index, err = bleve.Open("Movies")
		if err != nil {
			log.Panic(err)
		}
	}

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Panic(err)
	}
	err = cursor.All(ctx, &Movies)
	if err != nil {
		log.Panic(err)
	}

	for _, Movie := range Movies {
		Index.Index(Movie.ID.Hex(), Movie)
	}
}


func (Movie *Movie) AddIndex() {
	Index.Index(Movie.ID.Hex(), Movie)
}