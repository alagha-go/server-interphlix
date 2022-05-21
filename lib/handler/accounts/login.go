package accounts

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)


func LoginUrl(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	secretBody, err := ioutil.ReadFile("./secret1.json")
	HandlError(err)
	config, err := google.ConfigFromJSON(secretBody, "https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email")
	url := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	body := []byte(fmt.Sprintf(`{"login-url": "%s"}`, url))
	res.WriteHeader(200)
	res.Write(body)
}


func HandlError(err error) {
	if err != nil {
		log.Panic(err)
	}
}