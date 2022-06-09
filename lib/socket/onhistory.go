package socket

import (
	"encoding/json"
	"interphlix/lib/movies/history"
	"time"

	gosocketio "github.com/ambelovsky/gosf-socketio"
)


func OnHistory(channel *gosocketio.Channel, data string) interface{} {
	Channel, err := GetChannelByIP(channel.Ip())
	if err != nil {
		channel.Emit("error", `{"error": "client does not exist"}`)
		time.Sleep(500*time.Millisecond)
		channel.Close()
		return ""
	}
	var History history.History
	err = json.Unmarshal([]byte(data), &History)
	if err != nil {
		return `{"error": "invalid json"}`
	}
	History.AccountID = Channel.AccountID
	if History.Exists() {
		History.Update()
		return "done"
	}
	History.Create()
	return "done"
}