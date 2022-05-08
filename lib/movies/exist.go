package movies

import (
	"context"
	"interphlix/lib/variables"

	"go.mongodb.org/mongo-driver/bson"
)


func (movie *Movie) Exists() bool {
	var DBMovie Movie
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")

	err := collection.FindOne(ctx, bson.M{"code": movie.Code}).Decode(&DBMovie)
	if err != nil {
		return false
	}
	return true
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