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
	SocketServer = gosocketio.NewServer(transport.GetDefaultWebsocketTransport())
)

func StartSocketServer() {
	/// socket.io handlers
	SocketServer.On(gosocketio.OnConnection, OnConnection)
	SocketServer.On(gosocketio.OnDisconnection, OnDisconnection)
	SocketServer.On("online-users", GetOnlineUsers)
	SocketServer.On("login-url", GetUrl)
	SocketServer.On("rate-movie", OnRateMovie)
	SocketServer.On("rate-update", OnRateUpdate)

	serveMux := http.NewServeMux()
	serveMux.Handle("/socket.io/", SocketServer)

	log.Println("Starting SocketServer...")
	log.Panic(http.ListenAndServe(PORT, serveMux))
}



/// function to handle soicket.io's first connection
func OnConnection(channel *gosocketio.Channel) {
	var authorizationToken string
	Channel := Channel{ID: channel.Id(), IP: channel.Ip(), Channel: channel, TimeConnected: time.Now()}
	if len(channel.Request().Header["Cookie"]) > 0 {
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
	Channel.ServerID = Server.ID
	err := Channel.AddTODB()
	if err != nil {
		channel.Close()
	}
}


/// func to handle socket.io disconnection 
func OnDisconnection(channel *gosocketio.Channel) {
	DeleteChannel(channel.Id())
}