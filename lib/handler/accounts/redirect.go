package accounts

import (
	"context"
	"encoding/json"
	"fmt"
	"interphlix/lib/accounts"
	"interphlix/lib/socket"
	"interphlix/lib/variables"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)



func LoginRedirect(res http.ResponseWriter, req *http.Request) {
	var account accounts.Account
	code := req.URL.Query().Get("code")
	token := GetToken(code)
	json.Unmarshal(GetUserInfo(token.AccessToken), &account)
	account.Token = token
	data, _ := accounts.CreateAccount(account)
	json.Unmarshal(data, &account)
	tokenString, status, err := GenerateToken(account)
	if err != nil {
		res.WriteHeader(status)
		res.Write(variables.JsonMarshal(variables.Error{Error: err.Error()}))
	}
	cookie := &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Path: "/",
		Expires: time.Now().Add(120*time.Hour),
	}
	channel, err := socket.FindChannelByIP(req.Header["X-Real-Ip"][0])
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte("<h1>could not find a client application with your ip addres</h1>"))
		return
	}
	socket.EmitToken(channel, cookie)
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(fmt.Sprintf("<h1>Successfully loged in to Interphlix. Welcome %s</h1>", account.Name)))
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