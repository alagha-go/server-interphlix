package socket

import (
	"interphlix/lib/variables"
	"net/http"
	"time"

	gosocketio "github.com/ambelovsky/gosf-socketio"
)

/// socket.io function to get Authorization token
func EmitToken(channel *gosocketio.Channel, cookie *http.Cookie){
	var Channel Channel
	data := string(variables.JsonMarshal(cookie))
	channel.Emit("token", data)
	Account, _ := GetAccount(cookie.Value)
	DBChannel, err := GetChannelByID(Account.ID)
	if err == nil {
		channel.Emit("online", string(variables.JsonMarshal(DBChannel)))
		time.Sleep(500*time.Millisecond)
		if channel.Ip() != DBChannel.IP {
			channel.Close()
		}
		return
	}
	Channel.AccountID = Account.ID
	Channel.Channel = channel
	Channel.ID = channel.Id()
	Channel.IP = channel.Ip()
	Channel.TimeConnected = time.Now()
	Channel.ServerID = Server.ID
	DeleteChannel(channel.Id())
	Channel.AddTODB()
}