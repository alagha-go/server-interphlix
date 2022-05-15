package variables

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


type Secret struct {
	LocalDBUrl								string								`json:"local-db-url,omitempty"`
	RemoteDBUrl								string								`json:"remote-db-url,omitempty"`
}


var (
	Client	*mongo.Client
	Errors []Log
)

/// loads secret data from the the secret.json file
func LoadSecret() Secret {
	var secret Secret
	data, err := ioutil.ReadFile("./secret.json")
	HandleError(err, "variables","LoadSecret", "error while reading the sect.json file")
	json.Unmarshal(data, &secret)
	return secret
}


/// handle error by saving it to the DB and returning err == nil
func HandleError(err error, Package, function, comment string) bool {
	var Err bool = false
	if err != nil {
		Err = true
		var Log Log
		Log.Error = Error{Error: err.Error()}
		Log.Package = Package
		Log.Comment = comment
		Log.Function = function
		Log.HandleError()
	}
	return Err
}

/// function to marshal data to json
func JsonMarshal(data interface{}) []byte {
	var buff bytes.Buffer
	encoder := json.NewEncoder(&buff)
    encoder.SetEscapeHTML(false)
    encoder.Encode(data)
	return buff.Bytes()
}


func LoadErrors() {
	ctx := context.Background()
	collection := Client.Database("Interphlix").Collection("Errors")

	cursor, err := collection.Find(ctx, bson.M{})
	HandleError(err, "variables", "LoadErrors", "error while getting errors from the database")
	cursor.All(ctx, Errors)
}