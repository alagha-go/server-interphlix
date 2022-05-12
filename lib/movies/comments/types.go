package comments

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Comment struct {
	ID										primitive.ObjectID							`json:"_id,omitempty" bson:"_id,omitempty"`
	MovieID									primitive.ObjectID							`json:"movie-id,omitempty" bson:"movie-id,omitempty"`
	AccountID								primitive.ObjectID							`json:"account-id,omitempty" bson:"account-id,omitempty"`
	UserID									primitive.ObjectID							`json:"user-id,omitempty" bson:"user-id,omitempty"`
}

type CommentData struct {
	Comment									string										`json:"comment,omitempty" bson:"comment,omitempty"`
	TimeCommented							time.Time									`json:"time-commented,omitempty" bson:"time-commented,omitempty"`
}


type Reply struct {
	Reply									string										`json:"reply,omitempty" bson:"reply,omitempty"`
	TimeReplied								time.Time									`json:"time-replied,omitempty" bson:"time-replied,omitempty"`
}