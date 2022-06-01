package socket

import (
	"io/ioutil"
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

/// func to get google config
func GetConfig() (*oauth2.Config, error) {
	secretBody, err := ioutil.ReadFile("./secret1.json")
	HandlError(err)
	return google.ConfigFromJSON(secretBody, scopes...)
}

func HandlError(err error) {
	if err != nil {
		log.Panic(err)
	}
}