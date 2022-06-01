package socket

import (
	"fmt"
	"interphlix/lib/variables"
	"net/http"

	gosocketio "github.com/ambelovsky/gosf-socketio"
	"golang.org/x/oauth2"
)

/// socket.io function to get login url
func GetUrl(channel *gosocketio.Channel) interface{} {
	config, err := GetConfig()
	HandlError(err)
	url := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	return fmt.Sprintf(`{"url": "%s"}`, url)
}


/// socket.io function to get Authorization token
func GetToken(channel *gosocketio.Channel, cookie *http.Cookie){
	data := string(variables.JsonMarshal(cookie))
	channel.Emit("token", data)
}