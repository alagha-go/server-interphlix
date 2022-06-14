package movies

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
)


func GetTrendingMovies(seed int64) []Movie {
	var Movies []Movie
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Movies")

	cursor, _ := collection.Find(ctx, bson.M{"trending": true})
	cursor.All(ctx, &Movies)

	Movies = RandomMovies(seed, Movies)

	if len(Movies) > 5 {
		return Movies[:5]
	}

	return Movies
}