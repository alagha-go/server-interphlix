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

func (Type *Type) AddToLocalDB() {
	if TypeExists(Type.Type) {
		return
	}
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Types")

	_, err := collection.InsertOne(ctx, Type)
	variables.HandleError(err, "types", "type.AddToLocalDB", "error while inserting type to the local database")
}