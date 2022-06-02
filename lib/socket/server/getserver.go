package server

import (
	"context"
	"errors"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
)


func GetAllServers() ([]Server, error) {
	var Servers []Server
	ctx := context.Background()
	collection := variables.Client2.Database("Interphlix").Collection("Servers")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		variables.HandleError(err, "server", "GetAllServers", "error while getting servers from the database")
		return Servers, errors.New("could not get servers from the database")
	}
	err = cursor.All(ctx, &Servers)
	if err != nil {
		variables.HandleError(err, "server", "GetAllServers", "error while decoding cursor")
		return Servers, errors.New("could not decode cursor to servers")
	}
	return Servers, nil
}