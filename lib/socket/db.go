package socket

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
)


func (Channel *Channel) AddTODB() error {
	ctx := context.Background()
	collection := variables.Client2.Database("Interphlix").Collection("Channels")

	_, err := collection.InsertOne(ctx, Channel)
	return err
}

func DeleteChannel(id string) {
	ctx := context.Background()
	collection := variables.Client2.Database("Interphlix").Collection("Channels")

	_, err := collection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		DeleteChannel(id)
	}
}