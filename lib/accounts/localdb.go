package accounts

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
)

func (account *Account) AddToDB() {
	var Account Account
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Accounts")

	err := collection.FindOne(ctx, bson.M{"_id": account.ID}).Decode(&Account)
	if err != nil {
		account.Update()
		return
	}
	_, err = collection.InsertOne(ctx, account)
	variables.HandleError(err, "accounts", "aacount.AddToDB", "error while inserting account to the local database")
}