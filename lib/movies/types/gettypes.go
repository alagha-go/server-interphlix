package types

import (
	"context"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)


func GetTypes() ([]byte, int) {
	var Types []Type
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Types")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		variables.HandleError(err, "types", "GetTypes", "error while getting types from the database")
		return variables.JsonMarshal(variables.Error{Error: "could not get data from the database"}), http.StatusInternalServerError
	}
	cursor.All(ctx, &Types)
	return variables.JsonMarshal(Types), http.StatusOK
}