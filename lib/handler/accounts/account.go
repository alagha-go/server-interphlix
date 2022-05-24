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
	account, err := GetAccount(cookie.Value)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		res.Write(variables.JsonMarshal(variables.Error{Error: "account does not exist"}))
	}
	res.WriteHeader(http.StatusOK)
	res.Write(variables.JsonMarshal(account))
}