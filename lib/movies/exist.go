package movies

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/// check if a movie is in the Movies list
func (movie *Movie) Exists() bool {
	for _, Movie := range Movies {
		if movie.Code == Movie.Code && movie.Title == Movie.Title {
			return true
		}
	}
	return false
}


/// upload movie to the database
func (Movie *Movie) Upload() error {
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")
	Movie.SetServer()

	_, err := collection.InsertOne(ctx, Movie)
	if err != nil {
		variables.HandleError(err, "movies","Movie.Upload", "could not upload movie to the Database")
		return err
	}
	Movie.AddMovie()
	return nil
}


/// find a specific movie from the Movies list
func FindMovie(ID primitive.ObjectID) Movie {
	for _, Movie := range Movies {
		if Movie.ID == ID {
			return Movie
		}
	}
	return Movie{}
}

func (Movie *Movie) Valid() bool {
	return Movie.Title != ""
}