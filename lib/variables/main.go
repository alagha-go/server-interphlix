package variables

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"go.mongodb.org/mongo-driver/mongo"
)


type Secret struct {
	LocalDBUrl								string								`json:"local-db-url,omitempty"`
	RemoteDBUrl								string								`json:"remote-db-url,omitempty"`
}


var (
	Client	*mongo.Client
)

/// loads secret data from the the secret.json file
func LoadSecret() Secret {
	var secret Secret
	data, err := ioutil.ReadFile("./secret.json")
	HandleError(err, "LoadSecret", "error while reading the sect.json file")
	json.Unmarshal(data, &secret)
	return secret
}


/// handle error by saving it to the DB and returning err == nil
func HandleError(err error, function, comment string) bool {
	var Err bool = false
	if err != nil {
		Err = true
		var Log Log
		Log.Error = Error{Error: err.Error()}
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