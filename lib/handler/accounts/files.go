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
	account, err := GetAccount(cookie.Value)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		res.Write(variables.JsonMarshal(variables.Error{Error: "account does not exist"}))
	}
	data, status := drive.GetFiles(account)
	res.WriteHeader(status)
	res.Write(data)
}


func CreateFile(res http.ResponseWriter, req *http.Request) {
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
	name := req.URL.Query().Get("name")
	data, status := drive.CreateFile(account, name)
	res.WriteHeader(status)
	res.Write(data)
}


func DeleteFile(res http.ResponseWriter, req *http.Request) {
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
	id := req.URL.Query().Get("id")
	data, status := drive.DeleteFile(account, id)
	res.WriteHeader(status)
	res.Write(data)
}