package socket

import (
	"encoding/json"
	"fmt"
	"interphlix/lib/movies/ratings"
	"interphlix/lib/variables"
	"time"

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


/// socket.io function to handle ratemovie
func OnRateMovie(channel *gosocketio.Channel, data string) interface{} {
	Channel, err := GetChannelByIP(channel.Ip())
	if err != nil {
		channel.Emit("error", `{"error": "client does not exist"}`)
		time.Sleep(500*time.Millisecond)
		channel.Close()
		return ""
	}
	var Rate ratings.Rate
	err = json.Unmarshal([]byte(data), &Rate)
	if err != nil {
		return `{"error": "could not decode json data"}`
	}
	Rate.AccountID = Channel.AccountID
	content, _ := ratings.RateMovie(Rate)
	return string(content)
}