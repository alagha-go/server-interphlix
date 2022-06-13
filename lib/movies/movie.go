package movies

import (
	"context"
	"errors"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func GetMovieByID(MovieID primitive.ObjectID) (Movie, error) {
	var Movie Movie
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Movies")

	err := collection.FindOne(ctx, bson.M{"_id": MovieID}).Decode(&Movie)
	if err != nil {
		return Movie, errors.New("could not get Movie")
	}

	return Movie, nil
}