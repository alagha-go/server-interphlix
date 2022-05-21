package accounts

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
)



func LoginRedirect(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	code := req.URL.Query().Get("code")
	token := GetToken(code)
	res.WriteHeader(200)
	res.Write(GetUserInfo(token.AccessToken))
}


func GetUserInfo(token string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token=%s", token), nil)
	HandlError(err)
	res, err := client.Do(req)
	HandlError(err)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	HandlError(err)
	return body
}


func GetToken(code string) *oauth2.Token {
	config, err := GetConfig()
	HandlError(err)
	token, err := config.Exchange(context.Background(), code)
	HandlError(err)
	return token
}