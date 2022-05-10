package servers

import "go.mongodb.org/mongo-driver/bson/primitive"


type Server struct {
	ID										primitive.ObjectID							`json:"_id,omitempty" bson:"_id,omitempty"`
	Domain									string										`json:"domain,omitempty" bson:"domain,omitempty"`
	IPAddress								string										`json:"ip-address,omitempty" bson:"ip-address,omitempty"`
	Secure									bool										`json:"secure,omitempty" bson:"secure,omitempty"`
}