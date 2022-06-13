package casts

import (
	"context"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func SearchCasts(querry string, round int) ([]byte, int) {
	start := 0
	length := 50
	if round != 0 {
		if round == 1{
			start = 50
		}else {
			start = ((round-1)*20) + 50
		}
		length = (round*20) + 50
	}
	var Casts []Cast
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Casts")

	sort := bson.D{{"score", bson.D{{"$meta", "textScore"}}}}
	projection := bson.D{{"score", bson.D{{"$meta", "textScore"}}}}
	opts := options.Find().SetSort(sort).SetProjection(projection)

	cursor, err := collection.Find(ctx, bson.M{"$text": bson.M{"$search": querry}}, opts)
	if err != nil {
		return variables.JsonMarshal(variables.Error{Error: "could not search data"}), http.StatusInternalServerError
	}
	cursor.All(ctx, &Casts)

	if start > len(Casts) {
		return variables.JsonMarshal([]Cast{}), http.StatusOK
	}else if length > len(Casts) {
		return variables.JsonMarshal(Casts[start:]), http.StatusOK
	}

	return variables.JsonMarshal(Casts[start:length]), http.StatusOK
}