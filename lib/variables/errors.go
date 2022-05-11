package variables

import (
	"context"
	"time"

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

func GetErrors(Package string) ([]byte, int) {
	if Package == "" {
		return JsonMarshal(Errors), 200
	}
	var Logs []Log
	for _, Log := range Errors {
		if Log.Package == Package {
			Logs = append(Logs, Log)
		}
	}
	return JsonMarshal(Logs), 200
}