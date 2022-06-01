package socket

import (
	"log"
	"net/http"
	"time"

	gosocketio "github.com/ambelovsky/gosf-socketio"
	"github.com/ambelovsky/gosf-socketio/transport"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Channel struct {
	AccountID				primitive.ObjectID		`json:"_id,omitempty" bson:"_id,omitempty"`
	Channel					*gosocketio.Channel		`json:"channel,omitempty" bson:"channel,omitempty"`
	ID						string					`json:"id,omitempty" bson:"id,omitempty"`
	TimeConnected			time.Time				`json:"time_connected,omitempty" bson:"time_connected,omitempty"`
	IP						string					`json:"ip,omitempty" bson:"ip,omitempty"`
}

type Message struct {
	Id      int    `json:"id"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

var (
	PORT = ":9000"
	scopes = []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"}
)


func Main() {
	server := gosocketio.NewServer(transport.GetDefaultWebsocketTransport())

	/// socket.io handlers
	server.On(gosocketio.OnConnection, OnConnection)
	server.On(gosocketio.OnDisconnection, OnDisconnection)
	server.On("online-users", GetOnlineUsers)
	server.On("login-url", GetUrl)

	serveMux := http.NewServeMux()
	serveMux.Handle("/socket.io/", server)

	log.Println("Starting server...")
	log.Panic(http.ListenAndServe(PORT, serveMux))
}


/// function to handle soicket.io's first connection
func OnConnection(channel *gosocketio.Channel) {
	Channels = append(Channels, channel)
}


/// func to handle socket.io disconnection 
func OnDisconnection(channel *gosocketio.Channel) {
	for index := range Channels {
		if Channels[index].Id() == channel.Id() {
			RemoveChannel(index)
			return
		}
	}
}