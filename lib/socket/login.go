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
	Channels []*gosocketio.Channel
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


func FindChannelByIP(ip string) (*gosocketio.Channel, error) {
	for index := range Channels {
		if Channels[index].Ip() == ip {
			return Channels[index], nil
		}
	}
	return nil, errors.New("no client found")
}

func RemoveChannel(cahnnel *gosocketio.Channel, index int) {
	Channels[index] = Channels[len(Channels)-1]
    Channels = Channels[:len(Channels)-1]
}