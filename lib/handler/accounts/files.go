package accounts

import (
	"interphlix/lib/variables"
	"interphlix/lib/accounts/drive"
	"net/http"
)


func GetMyFiles(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	cookie, err := req.Cookie("token")
	if err != nil {
		res.WriteHeader(http.StatusUnauthorized)
		res.Write(variables.JsonMarshal(variables.Error{Error: "no account found for you"}))
		return
	}
	account := GetAccount(cookie.Value)
	data, status := drive.GetFiles(account)
	res.WriteHeader(status)
	res.Write(data)
}