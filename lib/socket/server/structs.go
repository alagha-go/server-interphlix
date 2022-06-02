package server

import "go.mongodb.org/mongo-driver/bson/primitive"


type Server struct {
	ID							primitive.ObjectID							`json:"_id,omitempty" bson:"_id,omitempty"`
	Domain						string										`json:"domain,omitempty" bson:"domain,omitempty"`
	IPAddress					string										`json:"ip_address,omitempty" bson:"ip_address,omitempty"`
	LocalIPAddress				string										`json:"local_ip_address,omitempty" bson:"local_ip_address,omitempty"`
	Secure						bool										`json:"secure,omitempty" bson:"secure,omitempty"`
	Working						bool										`json:"working,omitempty" bson:"working,omitempty"`
}