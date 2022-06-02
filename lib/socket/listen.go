package socket

import (
	"fmt"
	"interphlix/lib/variables"

	gosocketio "github.com/ambelovsky/gosf-socketio"
	"golang.org/x/oauth2"
)

/// socket.io function to get all online users
func GetOnlineUsers(channel *gosocketio.Channel) interface{} {
	Channels, err := GetAllChannels()
	if err != nil {
		return fmt.Sprintf(`{"error": "%s"`, err.Error())
	}
	data := variables.JsonMarshal(Channels)
	return string(data)
}


/// socket.io function to get login url
func GetUrl(channel *gosocketio.Channel) interface{} {
	config, err := GetConfig()
	HandlError(err)
	url := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	return fmt.Sprintf(`{"url": "%s"}`, url)
}