package movies

import (
	"encoding/json"
	"interphlix/lib/handler/accounts"
	"interphlix/lib/movies"
	"interphlix/lib/variables"
	"net/http"
	"strconv"
)

func GetMyHistory(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	err, status := accounts.ValidateRequest(req, "user")
	if err != nil {
		res.WriteHeader(status)
		res.Write(variables.JsonMarshal(variables.Error{Error: err.Error()}))
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
	round, err := strconv.Atoi(req.URL.Query().Get("round"))
	if err != nil {
		round = 0
	}
	data, status := movies.GetMyHistory(account.ID, round)
	res.WriteHeader(status)
	res.Write(data)
}


func GetContinue(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	err, status := accounts.ValidateRequest(req, "user")
	if err != nil {
		res.WriteHeader(status)
		res.Write(variables.JsonMarshal(variables.Error{Error: err.Error()}))
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
	round, err := strconv.Atoi(req.URL.Query().Get("round"))
	if err != nil {
		round = 1
	}
	start := round*20
	end := start+20
	Movies := movies.GetContinue(account.ID, start, end)
	json.NewEncoder(res).Encode(Movies)
}