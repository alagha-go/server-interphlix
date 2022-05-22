package accounts

import (
	"interphlix/lib/variables"
	"net/http"
)

func GetMyAccount(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	cookie, err := req.Cookie("token")
	if err != nil {
		res.WriteHeader(http.StatusUnauthorized)
		res.Write(variables.JsonMarshal(variables.Error{Error: "no account found for you"}))
		return
	}
	account := GetAccount(cookie.Value)
	res.WriteHeader(http.StatusOK)
	res.Write(variables.JsonMarshal(account))
}