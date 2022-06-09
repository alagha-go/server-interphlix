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
	account.Type = "user"
	_, err = collection.InsertOne(context.Background(), account)
	if err != nil {
		return variables.JsonMarshal(variables.Error{Error: "could not save user"}), http.StatusInternalServerError
	}
	return variables.JsonMarshal(account), http.StatusOK
}


func SetAccountStaff(ID primitive.ObjectID) ([]byte, int) {
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Accounts")

	filter := bson.M{"_id": bson.M{"$eq": ID}}
	update := bson.M{
		"$set": bson.M{
			"type": "staff",
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		variables.HandleError(err, "accounts", "SetAccountStaff", "error while updating document")
		return []byte(`{"error": "could not update account"}`), http.StatusInternalServerError
	}
	return []byte(`done`), http.StatusOK
}