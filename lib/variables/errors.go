package variables

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Log struct {
	ID									primitive.ObjectID				`json:"_id,omitempty" bson:"_id,omitempty"`
	Time								time.Time						`json:"time,omitempty" bson:"time,omitempty"`
	Error								Error							`json:"Error,omitempty" bson:"Error,omitempty"`
	Function							string							`json:"function,omitempty" bson:"function,omitempty"`
	Reason								string							`json:"reason,omitempty" bson:"reason,omitempty"`
}

type Error struct {
	Error								string							`json:"error,omitempty" bson:"error,omitempty"`
}


func (log *Log) HandleError() {
	ctx := context.Background()
	collection := LocalClient.Database("Interphlix").Collection("Errors")

	log.ID = primitive.NewObjectID()
	log.Time = time.Now()

	collection.InsertOne(ctx, log)
}