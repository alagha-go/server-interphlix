package socket

import (
	"interphlix/lib/variables"

	gosocketio "github.com/ambelovsky/gosf-socketio"
)


func GetOnlineUsers(channel *gosocketio.Channel) interface{} {
	users := variables.JsonMarshal(Channels)
	return string(users)
}