package types

import (
	"context"
	"interphlix/lib/variables"
)


func AddType(title string) {
	if TypeExists(title) {
		return
	}else if title == "" || title == " " {
		return
	}
	Type := Type{Type: title}
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Types")

	_, err := collection.InsertOne(ctx, Type)
	variables.HandleError(err, "types", "AddType", "error while inserting type to the remote database")
}


func (Type *Type) AddToDB() bool {
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Types")

	_, err := collection.InsertOne(ctx, Type)
	if err != nil {
		variables.HandleError(err, "types", "AddToDB", "error while adding type to the database")
		return false
	}
	return true
}

func (Type *Type) AddToLocalDB() {

}