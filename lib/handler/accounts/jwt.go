package accounts

import (
	"errors"
	"interphlix/lib/accounts"
	"interphlix/lib/variables"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Claims struct {
	AccountID					primitive.ObjectID
	jwt.StandardClaims
}

func GenerateToken(account accounts.Account) (string, int, error) {
	expires := time.Now().Add(120*time.Hour)
	claims := &Claims{
		AccountID: account.ID,
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


func VerifyToken(tokenString string) (bool, int) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		secret := variables.LoadSecret()
		return []byte(secret.JwtKey), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return false, http.StatusUnauthorized
		}
		return false, http.StatusBadRequest
	}

	if !token.Valid {
		return false, http.StatusUnauthorized
	}

	return true, http.StatusOK
}


func GetAccount(tokenString string) (accounts.Account, error) {
	claims := &Claims{}

	jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		secret := variables.LoadSecret()
		return []byte(secret.JwtKey), nil
	})
	account, err := accounts.GetAccount(claims.AccountID)
	return account, err
}


func RefreshToken(tokenString string) (string, int, error) {
	valid, status := VerifyToken(tokenString)
	if !valid {
		return "", status, errors.New("invalid token")
	}
	account, err := GetAccount(tokenString)
	if err != nil {
		return "", http.StatusNotFound, errors.New("account does not exist")
	}
	return GenerateToken(account)
}