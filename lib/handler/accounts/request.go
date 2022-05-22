package accounts

import (
	"net/http"
)


func ValidateRequest(req *http.Request) bool {
	cookie, err := req.Cookie("token")
	if err != nil {
		return false
	}
	valid, _ := VerifyToken(cookie.Value)
	return valid
}