package variables

import (
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)


func GetErrors(Package string) ([]byte, int) {
	var Errors []Log
	ctx := context.Background()
	collection := Client2.Database("Interphlix").Collection("Errors")
	if Package == "all" {
		cursor, err := collection.Find(ctx, bson.M{})
		HandleError(err, "variables", "GetErrors", "error while getting errors from the database")
		cursor.All(ctx, &Errors)
	}else {
		cursor, err := collection.Find(ctx, bson.M{"package": Package})
		HandleError(err, "variables", "GetErrors", "error while getting errors from the database")
		cursor.All(ctx, &Errors)
	}
	return JsonMarshal(Errors), http.StatusOK
}