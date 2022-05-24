package accounts

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
)


var (
	Accounts []Account
)

func Main() {
	LoadAccounts()
}

func LoadAccounts() {
	collection := variables.Client.Database("Interphlix").Collection("Accounts")

	cursor, err := collection.Find(context.Background(), bson.M{})
	variables.HandleError(err, "accounts", "LoadAccounts", "error while loading accounts from the database")
	cursor.All(context.Background(), &Accounts)
}