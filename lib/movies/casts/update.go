package casts

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
)


func SetImageUrl(name, url string) error {
	collection := variables.Client.Database("Interphlix").Collection("Casts")

	filter := bson.M{
		"name": name,
	}
	update := bson.M{"$set": bson.M{
		"image": url,
	}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	return err
}


func (cast *Cast) Update() {
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Casts")

	filter := bson.M{
		"name": cast.Name,
	}

	update := bson.M{"$set": cast}

	collection.UpdateOne(ctx, filter, update)
}