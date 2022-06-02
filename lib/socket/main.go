package socket

import (
	"interphlix/lib/variables"
	"log"
	"net/http"
	"time"

	gosocketio "github.com/ambelovsky/gosf-socketio"
	"github.com/ambelovsky/gosf-socketio/transport"
)

var (
	PORT = ":9000"
	scopes = []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"}
	Server = gosocketio.NewServer(transport.GetDefaultWebsocketTransport())
)


func Main() {

	/// socket.io handlers
	Server.On(gosocketio.OnConnection, OnConnection)
	Server.On(gosocketio.OnDisconnection, OnDisconnection)
	Server.On("online-users", GetOnlineUsers)
	Server.On("login-url", GetUrl)

	serveMux := http.NewServeMux()
	serveMux.Handle("/socket.io/", Server)

	log.Println("Starting Server...")
	log.Panic(http.ListenAndServe(PORT, serveMux))
}


/// function to handle soicket.io's first connection
func OnConnection(channel *gosocketio.Channel) {
	var authorizationToken string
	Channel := Channel{ID: channel.Id(), IP: channel.Ip(), Channel: channel, TimeConnected: time.Now()}
	if len(channel.Request().Header["Cookie"][0]) > 0 {
		authorizationToken = channel.Request().Header["Cookie"][0]
		Account, err := GetAccount(authorizationToken)
		if err == nil {
			Channel.AccountID = Account.ID
			Channel.Verified = true
		}
		DBChannel, err := GetChannelByID(Account.ID)
		if err == nil {
			channel.Emit("online", string(variables.JsonMarshal(DBChannel)))
			time.Sleep(500*time.Millisecond)
			channel.Close()
			return
		}
	}
	err := Channel.AddTODB()
	if err != nil {
		channel.Close()
	}
}


/// func to handle socket.io disconnection 
func OnDisconnection(channel *gosocketio.Channel) {
	DeleteChannel(channel.Id())
}