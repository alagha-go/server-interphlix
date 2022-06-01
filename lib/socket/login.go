package socket

import (
	"fmt"

	gosocketio "github.com/ambelovsky/gosf-socketio"
	"golang.org/x/oauth2"
)

func GetUrl(channel *gosocketio.Channel) interface{} {
	config, err := GetConfig()
	HandlError(err)
	url := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	return []byte(fmt.Sprintf(`{"url": "%s"}`, url))
}