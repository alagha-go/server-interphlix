package movies

import (
	"context"
	"interphlix/lib/movies/genres"
	"interphlix/lib/variables"
	"math/rand"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)


func GetMovies(round int) ([]byte, int) {
	var Movies []Movie
	start := 0
	end := 30
	collection := variables.Client1.Database("Interphlix").Collection("Movies")
	cursor, err := collection.Find(context.Background(), bson.M{})
	variables.HandleError(err, "movies", "GetMovies", "error while getting movies from the local database")
	cursor.All(context.Background(), &Movies)
	if round != 0 {
		start = round * 30
		end = round * 30 + 30
	}

	if start >= len(Movies) {
		return []byte(`{"error": "end"}`), http.StatusOK
	}

	if end >= len(Movies){
		return variables.JsonMarshal(Movies[start:]), http.StatusOK
	}

	return variables.JsonMarshal(Movies[start:end]), http.StatusOK
}


func GetHome() ([]byte, int) {
	var Movies []Movie
	var Genres []genres.Genre
	var recommendation Recommendation
	recommendation.Seed = time.Now().UnixNano()
	ctx := context.Background()
	collection1 := variables.Client1.Database("Interphlix").Collection("Genres")
	collection := variables.Client1.Database("Interphlix").Collection("Movies")
	Categories := []Category{{Title: "Featured", Movies: RandomMovies(recommendation.Seed, FeaturedMovies)}, {Title: "Popular Movies", Movies: RandomMovies(recommendation.Seed, PopularMovies)}, {Title: "Popular Tvs", Movies: RandomMovies(recommendation.Seed, PopularTvShows)}}
	recommendation.Categories = append(recommendation.Categories, Categories...)

	cursor, err := collection1.Find(ctx, bson.M{})
	if err != nil {
		variables.HandleError(err, "movies", "GetHome", "error while getting genres")
		return variables.JsonMarshal(variables.Error{Error: "could not get movies"}), http.StatusInternalServerError
	}
	err = cursor.All(ctx, &Genres)
	if err != nil {
		variables.HandleError(err, "movies", "GetHome", "error while decoding cursor")
		return variables.JsonMarshal(variables.Error{Error: "could not get movies"}), http.StatusInternalServerError
	}
	
	for index := range Genres {
		var Category Category
		Category.Title = Genres[index].Title
		recommendation.Categories = append(recommendation.Categories, Category)
	}

	cursor, err = collection.Find(ctx, bson.M{})
	if err != nil {
		variables.HandleError(err, "movies", "GetHome", "error while getting movies")
		return variables.JsonMarshal(variables.Error{Error: "could not get movies"}), http.StatusInternalServerError
	}
	err = cursor.All(ctx, &Movies)
	if err != nil {
		variables.HandleError(err, "movies", "GetHome", "error while decoding cursor")
		return variables.JsonMarshal(variables.Error{Error: "could not get movies"}), http.StatusInternalServerError
	}

	for _, Movie := range Movies {
		for _, genre := range Movie.Genres {
			for index, category := range recommendation.Categories {
				if genre == category.Title {
					recommendation.Categories[index].Movies = append(recommendation.Categories[index].Movies, Movie)
				}
			}
		}
	}

	for index := range recommendation.Categories {
		if len(recommendation.Categories[index].Movies) > 20 {
			recommendation.Categories[index].Movies = recommendation.Categories[index].Movies[:20]
		}
	}

	return variables.JsonMarshal(recommendation), http.StatusOK
}


func RandomMovies(seed int64, Movies []Movie) []Movie {
	rand.Seed(seed)
	rand.Shuffle(len(Movies), func(i, j int) { Movies[i], Movies[j] = Movies[j], Movies[i] })
	return Movies
}