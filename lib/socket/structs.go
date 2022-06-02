package socket

import (
	"time"

	gosocketio "github.com/ambelovsky/gosf-socketio"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Channel struct {
	AccountID				primitive.ObjectID		`json:"_id,omitempty" bson:"_id,omitempty"`
	ServerID				primitive.ObjectID		`json:"server_id,omitempty" bson:"server_id,omitempty"`
	Channel					*gosocketio.Channel		`json:"channel,omitempty" bson:"channel,omitempty"`
	ID						string					`json:"id,omitempty" bson:"id,omitempty"`
	TimeConnected			time.Time				`json:"time_connected,omitempty" bson:"time_connected,omitempty"`
	IP						string					`json:"ip,omitempty" bson:"ip,omitempty"`
	Verified				bool					`json:"verified,omitempty" bson:"verified,omitempty"`
	Staff					bool					`json:"staff,omitempty" bson:"staff,omitempty"`
	Admin					bool					`json:"admin,omitempty" bson:"admin,omitempty"`
}

type Message struct {
	Id      int    `json:"id"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}


type Claims struct {
	AccountID					primitive.ObjectID
	jwt.StandardClaims
}