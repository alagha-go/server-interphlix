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
	var documents []interface{}
	collection1 := variables.Client.Database("Interphlix").Collection("Accounts")
	collection := variables.Client1.Database("Interphlix").Collection("Accounts")

	cursor, err := collection1.Find(context.Background(), bson.M{})
	variables.HandleError(err, "accounts", "LoadAccounts", "error while loading accounts from the database")
	cursor.All(context.Background(), &documents)
	collection.InsertMany(context.Background(), documents)
}