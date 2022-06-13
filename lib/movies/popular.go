package movies

import (
	"context"
	"interphlix/lib/variables"

	"github.com/gocolly/colly"
	"go.mongodb.org/mongo-driver/bson"
)

func SetPopularMovies() {
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")
	var Titles []string
	collector := colly.NewCollector()

	collector.OnHTML("tbody", func(element *colly.HTMLElement) {
		element.ForEach("tr", func(_ int, element *colly.HTMLElement) {
			Titles = append(Titles, element.ChildText(".text-bold.text-large1"))
		})
	})

	collector.Visit("https://imdb-api.com/most-popular-movies")

	for index := range Titles {
		movie := Movie{Title: Titles[index]}
		exists, Movie := movie.ExistByTitle()
		if exists && !Movie.Popular {
			filter := bson.M{"_id": bson.M{"$eq": Movie.ID}}
			update := bson.M{"$set": bson.M{"popular": true}}
			collection.UpdateOne(ctx, filter, update)
		}
	}
}


func SetPopularTvShows() {
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")
	var Titles []string
	collector := colly.NewCollector()

	collector.OnHTML("tbody", func(element *colly.HTMLElement) {
		element.ForEach("tr", func(_ int, element *colly.HTMLElement) {
			Titles = append(Titles, element.ChildText(".text-bold.text-large1"))
		})
	})

	collector.Visit("https://imdb-api.com/most-popular-tvs")

	for index := range Titles {
		movie := Movie{Title: Titles[index]}
		exists, Movie := movie.ExistByTitle()
		if exists && !Movie.Popular {
			filter := bson.M{"_id": bson.M{"$eq": Movie.ID}}
			update := bson.M{"$set": bson.M{"popular": true}}
			collection.UpdateOne(ctx, filter, update)
		}
	}
}


func GetPopularMovies() ([]Movie, error) {
	var Movies []Movie
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Movies")

	cursor, err := collection.Find(ctx, bson.M{"type": "Movie", "popular": true})
	if err != nil {
		variables.HandleError(err, "movies", "GetPopularMovies", "error while getting data from the database")
	}
	cursor.All(ctx, &Movies)
	return Movies, nil
}


func GetPopularTvShows() ([]Movie, error) {
	var Movies []Movie
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Movies")

	cursor, err := collection.Find(ctx, bson.M{"type": "Tv-Show", "popular": true})
	if err != nil {
		variables.HandleError(err, "movies", "GetPopularMovies", "error while getting data from the database")
	}
	cursor.All(ctx, &Movies)
	return Movies, nil
}