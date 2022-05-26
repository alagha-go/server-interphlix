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


func (cast *Cast) AddToLocalDB() {
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Casts")
	if CastExists(cast.Name) {
		cast.Update()
		return
	}
	collection.InsertOne(ctx, cast)
}