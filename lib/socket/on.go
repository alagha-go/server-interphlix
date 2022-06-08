package socket

import (
	"interphlix/lib/movies/watchlist"
	"time"

	gosocketio "github.com/ambelovsky/gosf-socketio"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func CreateWatchlist(channel *gosocketio.Channel, id string) interface{} {
	Channel, err := GetChannelByIP(channel.Ip())
	if err != nil {
		channel.Emit("error", `{"error": "client does not exist"}`)
		time.Sleep(500*time.Millisecond)
		channel.Close()
		return ""
	}
	var WatchList watchlist.WatchList
	WatchList.MovieID, err = primitive.ObjectIDFromHex(id)
	if err != nil {
		return `{"error": "invalid id"}`
	}
	WatchList.AccountID = Channel.AccountID
	return WatchList.Create()
}


func DeleteWatchlist(channel *gosocketio.Channel, id string) interface{} {
	Channel, err := GetChannelByIP(channel.Ip())
	if err != nil {
		channel.Emit("error", `{"error": "client does not exist"}`)
		time.Sleep(500*time.Millisecond)
		channel.Close()
		return ""
	}
	var WatchList watchlist.WatchList
	WatchList.MovieID, err = primitive.ObjectIDFromHex(id)
	if err != nil {
		return `{"error": "invalid id"}`
	}
	WatchList.AccountID = Channel.AccountID
	return WatchList.Delete()
}