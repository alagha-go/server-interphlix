package movies

import (
	"context"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func GetMoviesByGenre(genre string, round int, seed int64) ([]byte, int) {
	var Movies []Movie
	var movies []Movie
	start := 0
	end := 30
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Movies")

	cursor, err := collection.Find(ctx, bson.M{})
	variables.HandleError(err, "movies", "GetMovieByGenre", "error while getting movies from the database")
	err = cursor.All(ctx, &Movies)
	variables.HandleError(err, "movies", "GetMovieByGenre", "error while decoding cursor")

	for _, movie := range Movies {
		if movie.ContainsGenre(genre) {
			movies = append(movies, Movie{ID: movie.ID, Code: movie.Code, Title: movie.Title, Type: movie.Type, ImageUrl: movie.ImageUrl})
		}
	}

	if seed > 0 {
		movies = RandomMovies(seed, movies)
	}

	if round != 0 {
		start = round * 20
		end = round * 20 + 20
	}

	if start >= len(movies) {
		return []byte(`{"error": "end"}`), http.StatusOK
	}

	if end >= len(movies) {
		return variables.JsonMarshal(movies[start:]), http.StatusOK
	}

	return variables.JsonMarshal(movies[start:end]), http.StatusOK
}


func GetMoviesByGenreAndType(Type, genre string, round int) ([]byte, int) {
	var Movies []Movie
	var movies []Movie
	start := 0
	end := 30
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Movies")

	cursor, err := collection.Find(ctx, bson.M{"type": Type})
	variables.HandleError(err, "movies", "GetMovieByGenre", "error while getting movies from the database")
	err = cursor.All(ctx, &Movies)
	variables.HandleError(err, "movies", "GetMovieByGenre", "error while decoding cursor")

	for _, movie := range Movies {
		if movie.ContainsGenre(genre) {
			movies = append(movies, Movie{ID: movie.ID, Code: movie.Code, Title: movie.Title, Type: movie.Type, ImageUrl: movie.ImageUrl})
		}
	}

	if round != 0 {
		start = round * 30
		end = round * 30 + 30
	}

	if start >= len(movies) {
		return []byte(`{"error": "end"}`), http.StatusOK
	}

	if end >= len(movies) {
		return variables.JsonMarshal(movies[start:]), http.StatusOK
	}

	return variables.JsonMarshal(movies[start:end]), http.StatusOK
}


func (Movie *Movie) ContainsGenre(genre string) bool {
	for _, Genre := range Movie.Genres {
		if Genre == genre {
			return true
		}
	}
	return false
}