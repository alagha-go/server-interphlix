package wtachlist

import "go.mongodb.org/mongo-driver/bson/primitive"


type WatchList struct {
	ID							primitive.ObjectID							`json:"_id,omitempty" bson:"_id,omitempty"`
	AccountID					primitive.ObjectID							`json:"account_id,omitempty" bson:"account_id,omitempty"`
	MovieID						primitive.ObjectID							`json:"movie_id,omitempty" bson:"movie_id,omitempty"`
}