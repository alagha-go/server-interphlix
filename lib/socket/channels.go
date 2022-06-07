package socket

import (
	"context"
	"interphlix/lib/socket/server"
	"interphlix/lib/variables"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)



func CleanChannelsCollection() {
	for {
		RemoveUnNeededChannels()
		time.Sleep(10*time.Second)
	}
}


func RemoveUnNeededChannels() {
	Servers, _ := server.GetAllServers()
	for _, Server := range Servers {
		if !Server.Working {
			ctx := context.Background()
			collection := variables.Client2.Database("Interphlix").Collection("Channels")

			collection.DeleteMany(ctx, bson.M{"server_id": Server.ID})
		}
	}
}


func RemoveMyChannels() {
	Servers, _ := server.GetAllServers()
	for _, Server := range Servers {
		if Server.LocalIPAddress == GetMyIPAddress() {
			ctx := context.Background()
			collection := variables.Client2.Database("Interphlix").Collection("Channels")

			collection.DeleteMany(ctx, bson.M{"server_id": Server.ID})
		}
	}
}