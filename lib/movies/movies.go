package movies

import (
	"context"
	"errors"
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

func (Movie *Movie) AddMovie() {
	Movies = append(Movies, *Movie)
}

func (Movie *Movie) GetIndex() (int, error) {
	for index, movie := range Movies {
		if movie.ID == Movie.ID {
			return index, nil
		}
	}
	return 0, errors.New("movie does not exist")
}