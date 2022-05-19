package movies

import (
	"context"
	"interphlix/lib/movies/genres"
	"interphlix/lib/movies/types"
	"interphlix/lib/variables"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


func UploadOneMovie(Movie Movie) ([]byte, int) {
	if strings.Contains(Movie.ID.Hex(), "00000000") {
		Movie.ID = primitive.NewObjectID()
	}
	if Movie.Exists() {
		return variables.JsonMarshal(Movie), http.StatusConflict
	}
	for _, genre := range Movie.Genres {
		var Genre genres.Genre
		Genre.ID = primitive.NewObjectID()
		Genre.Title = genre
		if !Genre.Exists() {
			Genre.Types = append(Genre.Types, Movie.Type)
			Genre.Upload()
		}else {
			Genre.Type = Movie.Type
			Genre.Update()
		}
	}
	err := Movie.Upload()
	if err != nil {
		return variables.JsonMarshal(variables.Error{Error: "could not save Movie to the Database"}), http.StatusInternalServerError
	}
	return variables.JsonMarshal(Movie), http.StatusCreated
}


/// upload movie to the database
func (Movie *Movie) Upload() error {
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")
	types.AddType(Movie.Type)

	_, err := collection.InsertOne(ctx, Movie)
	if err != nil {
		variables.HandleError(err, "movies","Movie.Upload", "could not upload movie to the Database")
		return err
	}
	return nil
}