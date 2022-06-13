package socket

import (
	"interphlix/lib/accounts"
	"interphlix/lib/movies/history"
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

func OnCode(channel *gosocketio.Channel, code string) interface{} {
	for index, conn := range Connections {
		if conn.Code == code {
			EmitToken(channel, conn.Cookie)
			RemoveIndex(index)
			return ""
		}
	}
	return "wrong code"
}

func RemoveIndex(index int){
	Connections[index] = Connections[len(Connections)-1] // Copy last element to index i.
	Connections = Connections[:len(Connections)-1]   // Truncate slice.
}

func OnWatch(channel *gosocketio.Channel, id string) interface{} {
	Channel, err := GetChannelByIP(channel.Ip())
	if err != nil {
		channel.Emit("error", `{"error": "client does not exist"}`)
		time.Sleep(500*time.Millisecond)
		channel.Close()
		return ""
	}
	Account, err := accounts.GetAccount(Channel.AccountID)
	if err != nil {
		return `{"error": "user does not exist"}`
	}
	if Account.Paid {
		return true
	}
	MovieID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return `{"error": "invalid id"}`
	}
	return history.MovieAllowed(Channel.AccountID, MovieID)
}