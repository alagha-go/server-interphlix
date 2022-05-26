package casts

import (
	"context"
	"encoding/json"
	"interphlix/lib/variables"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func Listener() {
	go ListenForCastsCollection()
}

func ListenForCastsCollection() {
	collection := variables.Client.Database("Interphlix").Collection("Casts")

	matchPipeline := bson.D{{"$match", bson.D{{"operationType", bson.D{{ "$in", bson.A{"insert", "update", "replace"} }}}}}}
	projectPipeline := bson.D{{ "$project", bson.D{{"fullDocument", 1}}}}

	opts := options.ChangeStream()
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
		var cast Cast
		json.Unmarshal(content, &cast)
		cast.AddToLocalDB()
	}
	ListenForCastsCollection()
}