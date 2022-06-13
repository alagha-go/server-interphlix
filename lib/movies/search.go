package movies

import (
	"context"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func SearchMovies(querry, Type, genre string, round int) ([]byte, int) {
	var length int = 50
	if round != 0 {
		length = (round*20) + 50
	}
	var Movies []Movie
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Movies")

	sort := bson.D{{"score", bson.D{{"$meta", "textScore"}}}}
	projection := bson.D{{"score", bson.D{{"$meta", "textScore"}}}}
	opts := options.Find().SetSort(sort).SetProjection(projection)

	if Type != "" {
		cursor, err := collection.Find(ctx, bson.M{"type": Type, "$text": bson.M{"$search": querry}}, opts)
		if err != nil {
			return variables.JsonMarshal(variables.Error{Error: "could not search data"}), http.StatusInternalServerError
		}
		cursor.All(ctx, &Movies)
	}else {
		cursor, err := collection.Find(ctx, bson.M{"$text": bson.M{"$search": querry}}, opts)
		if err != nil {
			return variables.JsonMarshal(variables.Error{Error: "could not search data"}), http.StatusInternalServerError
		}
		cursor.All(ctx, &Movies)
	}

	if len(Movies) < length {
		return variables.JsonMarshal(Movies), http.StatusOK
	}

	return variables.JsonMarshal(Movies[:length]), http.StatusOK
}