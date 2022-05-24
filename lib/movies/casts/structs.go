package casts

import "go.mongodb.org/mongo-driver/bson/primitive"


type Cast struct {
	ID									primitive.ObjectID							`json:"_id,omitempty" bson:"_id,omitempty"`
	Name								string										`json:"name,omitempty" bson:"name,omitempty"`
	Image								string										`json:"image,omitempty" bson:"image,omitempty"`
}