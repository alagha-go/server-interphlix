package movies

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
)


func LoadMovies() {
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")

	cursor, err := collection.Find(ctx, bson.M{})
	variables.HandleError(err, "LoadMovies", "error while getting movies from the Database")
	err = cursor.All(ctx, &Movies)
	variables.HandleError(err, "LoadMovies", "error while decoding movies from the cursor")
}