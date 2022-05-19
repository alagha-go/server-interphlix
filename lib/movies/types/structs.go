package types

import "go.mongodb.org/mongo-driver/bson/primitive"


type Type struct {
	ID											primitive.ObjectID						`json:"_id,omitempty" bson:"_id,omitempty"`
	Type										string									`json:"type,omitempty" bson:"type,omitempty"`
}