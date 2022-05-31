package socket

import (
	"log"
	"net/http"

	gosocketio "github.com/ambelovsky/gosf-socketio"
	"github.com/ambelovsky/gosf-socketio/transport"
)


type Channel struct {
	Channel string `json:"channel"`
}

type Message struct {
	Id      int    `json:"id"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

var (
	PORT = ":9000"
)


func Main() {
	server := gosocketio.NewServer(transport.GetDefaultWebsocketTransport())

	server.On(gosocketio.OnConnection, func(channel *gosocketio.Channel) {
		println("connected: ", channel.Id())
	})
	server.On(gosocketio.OnDisconnection, func(c *gosocketio.Channel) {
		log.Println("Disconnected")
	})

	serveMux := http.NewServeMux()
	serveMux.Handle("/socket.io/", server)

	log.Println("Starting server...")
	log.Panic(http.ListenAndServe(PORT, serveMux))
}