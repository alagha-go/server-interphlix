package genres

import "go.mongodb.org/mongo-driver/bson/primitive"

type Genre struct {
	ID											primitive.ObjectID						`json:"_id,omitempty" bson:"_id,omitempty"`
	Title										string									`json:"title,omitempty" bson:"title,omitempty"`
	Types										[]string								`json:"types,omitempty" bson:"types,omitempty"`
	Type										string									`json:"type,omitempty" bson:"type,omitempty"`
}