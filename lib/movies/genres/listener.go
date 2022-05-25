package genres

import (
	"context"
	"encoding/json"
	"errors"
	"interphlix/lib/variables"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Listener() {
	go ListenForGenresCollection()
}

func ListenForGenresCollection() {
	collection := variables.Client.Database("Interphlix").Collection("Genres")

	matchPipeline := bson.D{{"$match", bson.D{{"operationType", bson.D{{ "$in", bson.A{"insert", "update", "replace"} }}}}}}
	projectPipeline := bson.D{{ "$project", bson.D{{"fullDocument", 1}}}}

	opts := options.ChangeStream().SetMaxAwaitTime(2 * time.Second)
	opts.SetFullDocument("updateLookup")

	stream, err := collection.Watch(context.TODO(), mongo.Pipeline{matchPipeline, projectPipeline}, opts)
	if err != nil {
		log.Panic(err)
	}

	for stream.Next(context.TODO()) {
		var data bson.M
		if err := stream.Decode(&data); err != nil {
			panic(err)
		}
		content := variables.JsonMarshal(data["fullDocument"])
		var genre Genre
		json.Unmarshal(content, &genre)
		index, err := genre.GetIndex()
		if err != nil {
			Genres = append(Genres, genre)
		}
		Genres[index] = genre
	}
	ListenForGenresCollection()
}