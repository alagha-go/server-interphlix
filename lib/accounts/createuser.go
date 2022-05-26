package accounts

import (
	"context"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func CreateAccount(account Account) ([]byte, int) {
	var accountExist Account
	collection := variables.Client.Database("Interphlix").Collection("Accounts")

	err := collection.FindOne(context.Background(), bson.M{"id": account.GoogleID}).Decode(&accountExist)
	if err == nil {
		return variables.JsonMarshal(accountExist), http.StatusOK
	}
	account.ID = primitive.NewObjectID()
	_, err = collection.InsertOne(context.Background(), account)
	if err != nil {
		return variables.JsonMarshal(variables.Error{Error: "could not save user"}), http.StatusInternalServerError
	}
	return variables.JsonMarshal(account), http.StatusOK
}