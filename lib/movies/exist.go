package movies

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


func (movie *Movie) Exists() bool {
	for _, Movie := range Movies {
		if movie.ID == Movie.ID {
			return true
		}
	}
	return false
}


func (Movie *Movie) Upload() error {
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")

	_, err := collection.InsertOne(ctx, Movie)
	if err != nil {
		variables.HandleError(err, "Movie.Upload", "could not upload movie to the remoteDB")
		return err
	}
	Movie.AddMovie()
	return nil
}

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