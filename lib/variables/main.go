package variables

import (
	"bytes"
	"encoding/json"

	"go.mongodb.org/mongo-driver/mongo"
)


var (
	Client	*mongo.Client
	Client1 *mongo.Client
	Client2 *mongo.Client
	Errors []Log
)


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