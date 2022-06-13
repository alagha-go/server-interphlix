package movies

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
)


func GetTrendingMovie() []Movie {
	var Movies []Movie
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Movies")

	cursor, _ := collection.Find(ctx, bson.M{"trending": true})
	cursor.All(ctx, &Movies)
	for index := range Movies {
		Movies[index] = Movie{ID: Movies[index].ID, Code: Movies[index].Code, Title: Movies[index].Title, Type: Movies[index].Type, ImageUrl: Movies[index].ImageUrl}
	}

	return Movies
}