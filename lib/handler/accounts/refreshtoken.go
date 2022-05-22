package accounts

import (
	"interphlix/lib/variables"
	"net/http"
	"time"
)


func RenewToken(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	cookie, err := req.Cookie("token")
	if err != nil {
		res.WriteHeader(http.StatusUnauthorized)
		res.Write(variables.JsonMarshal(variables.Error{Error: "no account found for you"}))
		return
	}
	tokenString, status, err := RefreshToken(cookie.Value)
	if err != nil {
		res.WriteHeader(status)
		res.Write(variables.JsonMarshal(variables.Error{Error: err.Error()}))
		return
	}
	http.SetCookie(res, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: time.Now().Add(120*time.Hour),
	})
	res.WriteHeader(status)
	res.Write([]byte(`{"success": true}`))
}