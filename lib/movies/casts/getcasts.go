package casts

import (
	"context"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAllCasts() ([]byte, int) {
	var Casts []Cast
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Casts")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		variables.HandleError(err, "catst", "GetAllCasts", "error while getting casts from the local database")
		return variables.JsonMarshal(variables.Error{Error: "could not get casts from the database"}), http.StatusInternalServerError
	}
	cursor.All(ctx, &Casts)
	return variables.JsonMarshal(Casts), http.StatusOK
}