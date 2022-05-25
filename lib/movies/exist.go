package movies

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/// check if a movie is in the Movies list
func (movie *Movie) Exists() bool {
	var Movie Movie
	collection := variables.Client1.Database("Interphlix").Collection("Movies")

	err := collection.FindOne(context.Background(), bson.M{"code": movie.Code}).Decode(&Movie)
	if err != nil {
		return false
	}
	return true
}



/// find a specific movie from the Movies list
func FindMovie(ID primitive.ObjectID) Movie {
	var Movie Movie
	collection := variables.Client1.Database("Interphlix").Collection("Movies")

	collection.FindOne(context.Background(), bson.M{"_id": ID}).Decode(&Movie)
	return Movie
}

func (Movie *Movie) Valid() bool {
	return Movie.Title != ""
}