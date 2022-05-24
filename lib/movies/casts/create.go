package casts

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


func CreateCast(name string) {
	Cast := Cast{ID: primitive.NewObjectID() ,Name: name}
	collection := variables.Client.Database("Interphlix").Collection("Casts")
	if CastExists(name) {
		return
	}
	
	collection.InsertOne(context.Background(), Cast)
}