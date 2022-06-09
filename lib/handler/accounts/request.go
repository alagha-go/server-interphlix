package accounts

import (
	"errors"
	"net/http"
)


func ValidateRequest(req *http.Request, Type string) (error, int) {
	cookie, err := req.Cookie("token")
	if err != nil {
		return errors.New("provide authorization token"), http.StatusUnauthorized
	}
	valid, _ := VerifyToken(cookie.Value)
	if !valid {
		return errors.New("invalid token"), http.StatusUnauthorized
	}
	account, err := GetAccount(cookie.Value)
	if err != nil {
		return err, http.StatusNotFound
	}
	if Type == account.Type {
		return nil, 200
	}else if Type == "staff" && account.Type != "user" {
		return nil, 200
	}else if Type == "user" {
		return nil, 200
	}else {
		return errors.New("forbidden"), http.StatusForbidden
	}
}