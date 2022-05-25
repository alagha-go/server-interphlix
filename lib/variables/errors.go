package variables

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Log struct {
	ID									primitive.ObjectID				`json:"_id,omitempty" bson:"_id,omitempty"`
	Time								time.Time						`json:"time,omitempty" bson:"time,omitempty"`
	Package								string							`json:"package,omitempty" bson:"package,omitempty"`
	Error								Error							`json:"Error,omitempty" bson:"Error,omitempty"`
	Function							string							`json:"function,omitempty" bson:"function,omitempty"`
	Comment								string							`json:"comment,omitempty" bson:"comment,omitempty"`
}

type Error struct {
	Error								string							`json:"error,omitempty" bson:"error,omitempty"`
}

/// handle error by saving to the database for feature reference
func (log *Log) HandleError() {
	ctx := context.Background()
	collection := Client.Database("Interphlix").Collection("Errors")

	log.ID = primitive.NewObjectID()
	log.Time = time.Now()

	_, err := collection.InsertOne(ctx, log)
	if err == nil {
		Errors = append(Errors, *log)
	}
}


func LoadErrors() {
	var Errors []interface{}
	ctx := context.Background()
	collection1 := Client2.Database("Interphlix").Collection("Errors")
	collection := Client.Database("Interphlix").Collection("Errors")

	cursor, err := collection1.Find(ctx, bson.M{})
	HandleError(err, "variables", "LoadErrors", "error while getting errors from the database")
	cursor.All(ctx, &Errors)
	collection.Drop(ctx)
	collection.InsertMany(ctx, Errors)
}