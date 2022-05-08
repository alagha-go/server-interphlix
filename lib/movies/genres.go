package movies

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
)


func LoadGenres() {
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Genres")

	cursor, err := collection.Find(ctx, bson.M{})
	variables.HandleError(err, "LoadGenres", "error while getting genres from the Database")
	err = cursor.All(ctx, &Genres)
	variables.HandleError(err, "LoadGenres", "error while decoding genres from the cursor")
}