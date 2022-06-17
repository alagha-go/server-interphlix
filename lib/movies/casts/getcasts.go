package casts

import (
	"context"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAllCasts(round int) ([]byte, int) {
	start := round * 30
	end := start + 30
	var Casts []Cast
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Casts")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		variables.HandleError(err, "catst", "GetAllCasts", "error while getting casts from the local database")
		return variables.JsonMarshal(variables.Error{Error: "could not get casts from the database"}), http.StatusInternalServerError
	}
	cursor.All(ctx, &Casts)

	if start > len(Casts) {
		return variables.JsonMarshal([]Cast{}), http.StatusOK
	}else if end > len(Casts) {
		return variables.JsonMarshal(Casts[start:]), http.StatusOK
	}
	return variables.JsonMarshal(Casts[start:end]), http.StatusOK
}