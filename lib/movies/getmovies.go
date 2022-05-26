package movies

import (
	"context"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func GetMoviesByGenre(genre string) ([]byte, int) {
	var Movies []Movie
	var movies []Movie
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Movies")

	cursor, err := collection.Find(ctx, bson.M{"genres": bson.M{"$all": bson.A{genre}}})
	variables.HandleError(err, "movies", "GetMovieByGenre", "error while getting movies from the database")
	err = cursor.All(ctx, &Movies)
	variables.HandleError(err, "movies", "GetMovieByGenre", "error while decoding cursor")

	for _, Movie := range Movies {
		if Movie.ContainsGenre(genre) {
			movies = append(movies, Movie)
		}
	}

	return variables.JsonMarshal(movies), http.StatusOK
}


func GetMoviesByGenreAndType(Type, genre string) ([]byte, int) {
	var Movies []Movie
	var movies []Movie
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Movies")

	cursor, err := collection.Find(ctx, bson.M{"type": Type})
	variables.HandleError(err, "movies", "GetMovieByGenre", "error while getting movies from the database")
	err = cursor.All(ctx, &Movies)
	variables.HandleError(err, "movies", "GetMovieByGenre", "error while decoding cursor")

	for _, Movie := range Movies {
		if Movie.ContainsGenre(genre) {
			movies = append(movies, Movie)
		}
	}

	return variables.JsonMarshal(movies), http.StatusOK
}


func (Movie *Movie) ContainsGenre(genre string) bool {
	for _, Genre := range Movie.Genres {
		if Genre == genre {
			return true
		}
	}
	return false
}