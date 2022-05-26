package accounts

import (
	"context"
	"errors"
	"interphlix/lib/variables"
	"net/http"

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



func GetAccounts() ([]byte, int) {
	var Accounts []Account
	collection := variables.Client1.Database("Interphlix").Collection("Accounts")
	ctx := context.Background()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		variables.HandleError(err, "accounts", "GetAccounts", "error while getting data from the local database")
		return variables.JsonMarshal(variables.Error{Error: "could not get accounts"}), http.StatusInternalServerError
	}
	cursor.All(ctx, &Accounts)
	return variables.JsonMarshal(Accounts), http.StatusOK
}