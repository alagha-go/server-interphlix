package accounts

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


func GetAccount(ID primitive.ObjectID) (Account, error) {
	for _, account := range Accounts {
		if account.ID == ID {
			if account.Deleted {
				return Account{}, errors.New("account does not exist")
			}else {
				return account, nil
			}
		}
	}
	return Account{}, errors.New("account does not exist")
}