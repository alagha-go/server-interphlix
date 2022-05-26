package casts

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
)


func CastExists(name string) bool {
	var Cast Cast
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Casts")

	err := collection.FindOne(ctx, bson.M{"name": name}).Decode(&Cast)
	return err == nil
}