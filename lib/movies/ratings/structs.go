package ratings

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Rate struct {
	ID									primitive.ObjectID								`json:"_id,omitempty" bson:"_id,omitempty"`
	AccountID							primitive.ObjectID								`json:"account-id,omitempty" bson:"account-id,omitempty"`
	UserID								primitive.ObjectID								`json:"user-id,omitempty" bson:"user-id,omitempty"`
	Stars								int												`json:"stars,omitempty" bson:"stars,omitempty"`
	Review								string											`json:"review,omitempty" bson:"review,omitempty"`
	TimeRated							time.Time										`json:"time-rated,omitempty" bson:"time-rated,omitempty"`
}