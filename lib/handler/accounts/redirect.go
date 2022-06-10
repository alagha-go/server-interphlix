package accounts

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"interphlix/lib/accounts"
	"interphlix/lib/socket"
	"interphlix/lib/variables"
	"io/ioutil"
	"math"
	"math/big"
	"net/http"
	"strconv"
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
	code = GenerateCode(4)
	socket.Connections = append(socket.Connections, socket.Connection{Code: code, Cookie: cookie})
	http.SetCookie(res, cookie)
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(fmt.Sprintf("<h1>G-%s</h1>", code)))
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


func GenerateCode(length int) string {
	var code string
	maxLimit := int64(int(math.Pow10(length)) - 1)
	lowLimit := int(math.Pow10(length - 1))

	randomNumber, _ := rand.Int(rand.Reader, big.NewInt(maxLimit))
	randomNumberInt := int(randomNumber.Int64())

	// Handling integers between 0, 10^(n-1) .. for n=4, handling cases between (0, 999)
	if randomNumberInt <= lowLimit {
		randomNumberInt += lowLimit
	}

	// Never likely to occur, kust for safe side.
	if randomNumberInt > int(maxLimit) {
		randomNumberInt = int(maxLimit)
	}

	for _, conn := range socket.Connections {
		if conn.Code == code {
			GenerateCode(length)
		}
	}

	code = strconv.Itoa(randomNumberInt)

	return code
}