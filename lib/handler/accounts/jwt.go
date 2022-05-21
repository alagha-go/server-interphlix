package accounts

import (
	"errors"
	"interphlix/lib/accounts"
	"interphlix/lib/variables"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Account					accounts.Account
	jwt.StandardClaims
}

func GenerateToken(account accounts.Account) (string, int, error) {
	expires := time.Now().Add(120*time.Hour)
	claims := &Claims{
		Account: account,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expires.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := variables.LoadSecret()
	tokenString, err := token.SignedString([]byte(secret.JwtKey))
	if err != nil {
		return "", http.StatusInternalServerError, errors.New("could not generate token")
	}
	return tokenString, 200, nil
}