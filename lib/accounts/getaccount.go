package accounts

import (
	"context"
	"errors"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func GetAccount(ID primitive.ObjectID) (Account, error) {
	var account Account
	collection := variables.Client1.Database("Interphlix").Collection("Accounts")
	ctx := context.Background()

	err := collection.FindOne(ctx, bson.M{"_id": ID}).Decode(&account)
	if err != nil {
		return account, errors.New("account not found")
	}
	return account, nil
}
