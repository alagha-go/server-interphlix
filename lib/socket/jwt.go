package socket

import (
	"interphlix/lib/accounts"
	"interphlix/lib/variables"

	"github.com/dgrijalva/jwt-go"
)

func GetAccount(tokenString string) (accounts.Account, error) {
	claims := &Claims{}

	jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		secret := variables.LoadSecret()
		return []byte(secret.JwtKey), nil
	})
	account, err := accounts.GetAccount(claims.AccountID)
	return account, err
}