package socket

import (
	"errors"
	"fmt"
	"interphlix/lib/variables"
	"net/http"

	gosocketio "github.com/ambelovsky/gosf-socketio"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

//// function to get channel by ip address
func FindChannelByIP(IP string) (*gosocketio.Channel, error) {
	Channel, err := GetChannelByIP(IP)
	if err != nil {
		return nil, err
	}
	return Server.GetChannel(Channel.ID)
}


/// check if account is online
func IsAccountOnline(ID primitive.ObjectID) bool {
	for index := range Channels {
		if Channels[index].AccountID == ID {
			return true
		}
	}
	return false
}


//// removes channel from the channels list
func RemoveChannel(index int) {
	Channels[index] = Channels[len(Channels)-1]
    Channels = Channels[:len(Channels)-1]
}

/// function to get channel by accountID

func GetChannel(AccountID primitive.ObjectID) *Channel {
	for index := range Channels {
		if Channels[index].AccountID == AccountID {
			return &Channels[index]
		}
	}
	return nil
}