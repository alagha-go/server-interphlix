package socket

import (
	"fmt"

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

//// function to get channel by ip address
func FindChannelByIP(IP string) (*gosocketio.Channel, error) {
	Channel, err := GetChannelByIP(IP)
	if err != nil {
		return nil, err
	}
	return Server.GetChannel(Channel.ID)
}