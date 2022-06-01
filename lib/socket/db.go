package socket

import (
	"context"
	"interphlix/lib/variables"
)


func (Channel *Channel) AddTODB() error {
	ctx := context.Background()
	collection := variables.Client2.Database("Interphlix").Collection("Channels")

	_, err := collection.InsertOne(ctx, Channel)
	return err
}