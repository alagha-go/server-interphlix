package accounts

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	scopes = []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/drive.file"}
)


func LoginUrl(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	config, err := GetConfig()
	HandlError(err)
	url := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	body := []byte(fmt.Sprintf(`{"login-url": "%s"}`, url))
	res.WriteHeader(200)
	res.Write(body)
}

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