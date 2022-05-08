package movies

import (
	"context"
	"interphlix/lib/variables"
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
	return nil
}