package socket

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/// function adds channel to the database
func (Channel *Channel) AddTODB() error {
	ctx := context.Background()
	collection := variables.Client2.Database("Interphlix").Collection("Channels")

	_, err := collection.InsertOne(ctx, Channel)
	return err
}


/// function deletes channel from the database
func DeleteChannel(id string) {
	ctx := context.Background()
	collection := variables.Client2.Database("Interphlix").Collection("Channels")

	_, err := collection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		DeleteChannel(id)
	}
}

//// function gets channel from the database
func GetChannelByID(ID primitive.ObjectID) (Channel, error) {
	var Channel Channel
	ctx := context.Background()
	collection := variables.Client2.Database("Interphlix").Collection("Channels")

	err := collection.FindOne(ctx, bson.M{"_id": ID}).Decode(&Channel)
	return Channel, err
}


//// function Gets channel from the database with ip filter
func GetChannelByIP(IP string) (Channel, error) {
	var Channel Channel
	ctx := context.Background()
	collection := variables.Client2.Database("Interphlix").Collection("Channels")

	err := collection.FindOne(ctx, bson.M{"ip": IP}).Decode(&Channel)
	return Channel, err
}