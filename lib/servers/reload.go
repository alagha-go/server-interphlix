package servers

import (
	"context"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)


func ReloadServers() {
	var Servers []Server
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Servers")

	cursor, err := collection.Find(ctx, bson.M{})
	variables.HandleError(err, "ReloadServers", "error while getting servers from the database")
	cursor.All(ctx, &Servers)

	for _, Server := range Servers {
		var protocol string = "http://"
		if Server.Secure {
			protocol = "https://"
		}
		url := protocol+Server.Domain+"/servers/reload"
		GetRequest(url)
	}

}

func GetRequest(url string) {
	http.Get(url)
}