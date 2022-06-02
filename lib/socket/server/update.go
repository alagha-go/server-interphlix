package server

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
)

func (Server *Server) SetWorking(working bool) error {
	ctx := context.Background()
	collection := variables.Client2.Database("Interphlix").Collection("Servers")

	filter := bson.M{
		"_id": bson.M{
			"$eq": Server.ID, // check if bool field has value of 'false'
		},
	}
	update := bson.M{"$set": bson.M{
		"working": working,
	}}

	_, err := collection.UpdateOne(ctx, filter, update)
	variables.HandleError(err, "server", "Server.SetWorking", "error while updating server")
	return err
}