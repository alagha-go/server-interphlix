package socket

import (
	"errors"
	"fmt"
	"interphlix/lib/variables"
	"net/http"

	gosocketio "github.com/ambelovsky/gosf-socketio"
	"golang.org/x/oauth2"
)

var (
	Channels []Channel
)

/// socket.io function to get login url
func GetUrl(channel *gosocketio.Channel) interface{} {
	config, err := GetConfig()
	HandlError(err)
	url := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	return fmt.Sprintf(`{"url": "%s"}`, url)
}


/// socket.io function to get Authorization token
func EmitToken(channel *gosocketio.Channel, cookie *http.Cookie){
	data := string(variables.JsonMarshal(cookie))
	channel.Emit("token", data)
}


func FindChannelByIP(IP string) (*gosocketio.Channel, error) {
	for index := range Channels {
		if Channels[index].IP == IP {
			return Server.GetChannel(Channels[index].ID)
		}
	}
	return nil, errors.New("no client found")
}

func RemoveChannel(index int) {
	Channels[index] = Channels[len(Channels)-1]
    Channels = Channels[:len(Channels)-1]
}