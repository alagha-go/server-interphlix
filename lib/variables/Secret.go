package variables

import (
	"encoding/json"
	"io/ioutil"
)

type Secret struct {
	LocalDBUrl								string								`json:"local-db-url,omitempty"`
	RemoteDBUrl								string								`json:"remote-db-url,omitempty"`
	Remote2DBUrl							string								`json:"remote2-db-url,omitempty"`
	JwtKey									string								`json:"jwtkey,omitempty"`
}

/// loads secret data from the the secret.json file
func LoadSecret() Secret {
	var secret Secret
	data, err := ioutil.ReadFile("./secret.json")
	HandleError(err, "variables","LoadSecret", "error while reading the sect.json file")
	json.Unmarshal(data, &secret)
	return secret
}