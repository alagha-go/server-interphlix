package types

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


func AddType(title string) {
	if TypeExists(title) {
		return
	}else if title == "" || title == " " {
		return
	}
	var Type Type
	Type.ID = primitive.NewObjectID()
	Type.Type = title
	inserted := Type.AddToDB()
	if inserted {
		Types = append(Types, Type)
	}
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