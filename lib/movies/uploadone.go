package movies

import (
	"interphlix/lib/movies/genres"
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
		if Movie.Type == "Tv-Show"{
			Genre.TvShow = true
		}else if Movie.Type == "Movie" {
			Genre.Movie = true
		}else if Movie.Type == "Fanproj" {
			Genre.Fanproj = true
		}else {
			Genre.Afro = true
		}
		if !Genre.Exists() {
			Genre.Upload()
		}else {
			Genre.Update()
		}
	}
	err := Movie.Upload()
	if err != nil {
		return variables.JsonMarshal(variables.Error{Error: "could not save Movie to the Database"}), http.StatusInternalServerError
	}
	return variables.JsonMarshal(Movie), http.StatusCreated
}