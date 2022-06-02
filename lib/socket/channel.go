package socket

import (
	"context"
	"errors"
	"interphlix/lib/variables"

	gosocketio "github.com/ambelovsky/gosf-socketio"
	"go.mongodb.org/mongo-driver/bson"
)

//// function to get channel by ip address
func FindChannelByIP(IP string) (*gosocketio.Channel, error) {
	Channel, err := GetChannelByIP(IP)
	if err != nil {
		return nil, err
	}
	return Server.GetChannel(Channel.ID)
}


func GetAllChannels() ([]Channel, error) {
	var Channels []Channel
	ctx := context.Background()
	collection := variables.Client2.Database("Interphlix").Collection("Channels")
	
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		variables.HandleError(err, "socket", "GetAllChannels", "error while getting chennles from the database")
		return Channels, errors.New("error could not get channels from the database")
	}
	err = cursor.All(ctx, &Channels)
	if err != nil {
		variables.HandleError(err, "socket", "GetAllChannels", "error while decoding cursor")
		return Channels, errors.New("error could not decode data")
	}
	return Channels, nil
}