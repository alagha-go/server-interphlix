package variables

import (
	"encoding/json"
	"io/ioutil"

	"go.mongodb.org/mongo-driver/mongo"
)


type Secret struct {
	LocalDBUrl								string								`json:"local-db-url,omitempty"`
	RemoteDBUrl								string								`json:"remote-db-url,omitempty"`
}


var (
	LocalClient	*mongo.Client
	RemoteClient *mongo.Client
)

func LoadSecret() Secret {
	var secret Secret
	data, err := ioutil.ReadFile("./secret.json")
	HandleError(err, "LoadSecret", "error while reading the sect.json file")
	json.Unmarshal(data, &secret)
	return secret
}

func HandleError(err error, function, reason string) bool {
	var Err bool
	if err != nil {
		Err = true
		var Log Log
		Log.Error = Error{Error: err.Error()}
		Log.Reason = reason
		Log.Function = function
		Log.HandleError()
	}
	return Err == true
}