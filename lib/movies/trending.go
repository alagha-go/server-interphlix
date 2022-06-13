package movies

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
)


func GetTrendingMovies() []Movie {
	var Movies []Movie
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Movies")

	cursor, _ := collection.Find(ctx, bson.M{"trending": true})
	cursor.All(ctx, &Movies)

	return Movies
}