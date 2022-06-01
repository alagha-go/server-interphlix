package accounts

import (
	"interphlix/lib/accounts"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)


func Main() {
	accounts.Main()
}

func HandlError(err error) {
	if err != nil {
		log.Panic(err)
	}
}


func GetConfig() (*oauth2.Config, error) {
	secretBody, err := ioutil.ReadFile("./secret1.json")
	HandlError(err)
	return google.ConfigFromJSON(secretBody, scopes...)
}