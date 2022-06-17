package movies

import (
	"context"
	"interphlix/lib/variables"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)


func GetTrendingMovies(seed int64) []Movie {
	var Movies []Movie
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Movies")

	cursor, _ := collection.Find(ctx, bson.M{"trending": true})
	cursor.All(ctx, &Movies)

	Movies = RandomMovies(seed, Movies)

	if len(Movies) > 5 {
		return Movies[:5]
	}

	return Movies
}


func GetTrendingMoviesApi(seed int64, round int) ([]byte, int) {
	start := 0
	end := 5
	if round != 0 {
		start = round * 5
		end = start+5
	}
	var Movies []Movie
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Movies")

	cursor, _ := collection.Find(ctx, bson.M{"trending": true})
	cursor.All(ctx, &Movies)

	Movies = RandomMovies(seed, Movies)

	for index, movie := range Movies {
		Movies[index] = Movie{ID: movie.ID, Code: movie.Code, Title: movie.Title, Type: movie.Type, ImageUrl: movie.ImageUrl}
	}

	if start > len(Movies) {
		return variables.JsonMarshal([]Movie{}), http.StatusOK
	}else if end > len(Movies) {
		return variables.JsonMarshal(Movies[start:]), http.StatusOK
	}

	return variables.JsonMarshal(Movies[start:end]), http.StatusOK
}