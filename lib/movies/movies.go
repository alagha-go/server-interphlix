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

func ListenMovies() {
	var movies []Movie
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")

	cursor, err := collection.Find(ctx, bson.M{})
	variables.HandleError(err, "ListenMovies", "error while listening for changes in movies collection")
	err = cursor.All(ctx, &movies)
	variables.HandleError(err, "ListenMovies", "error while decoding movies from the cursor")
	for _, movie := range movies {
		if movie.Exists() {
			for movieIndex, Movie := range Movies {
				if Movie.ID == movie.ID {
					Movies[movieIndex] = movie
				}
			}
		}else {
			Movies = append(Movies, movie)
		}
	}
	ListenMovies()
}