package movies

import (
	"interphlix/lib/handler/accounts"
	"interphlix/lib/movies"
	"interphlix/lib/variables"
	"net/http"
)

func GetRatedMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	valid := accounts.ValidateRequest(req)
	if !valid {
		res.WriteHeader(http.StatusUnauthorized)
		res.Write(variables.JsonMarshal(variables.Error{Error: "unauthorized"}))
		return
	}
	cookie, err := req.Cookie("token")
	if err != nil {
		res.WriteHeader(http.StatusUnauthorized)
		res.Write(variables.JsonMarshal(variables.Error{Error: "unauthorized"}))
		return
	}
	account, err := accounts.GetAccount(cookie.Value)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		res.Write(variables.JsonMarshal(variables.Error{Error: "account does not exist"}))
		return
	}
	data, status := movies.GetRatedMovies(account.ID)
	res.WriteHeader(status)
	res.Write(data)
}


func GetMyWatchlist(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	valid := accounts.ValidateRequest(req)
	if !valid {
		res.WriteHeader(http.StatusUnauthorized)
		res.Write(variables.JsonMarshal(variables.Error{Error: "unauthorized"}))
		return
	}
	cookie, err := req.Cookie("token")
	if err != nil {
		res.WriteHeader(http.StatusUnauthorized)
		res.Write(variables.JsonMarshal(variables.Error{Error: "unauthorized"}))
		return
	}
	account, err := accounts.GetAccount(cookie.Value)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		res.Write(variables.JsonMarshal(variables.Error{Error: "account does not exist"}))
		return
	}
	data, status := movies.GetMyWatchlist(account.ID)
	res.WriteHeader(status)
	res.Write(data)
}