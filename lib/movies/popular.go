package movies

import (
	"context"
	"interphlix/lib/variables"

	"github.com/gocolly/colly"
	"go.mongodb.org/mongo-driver/bson"
)

func GetPopularMovies() {
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Movies")
	var Titles []string
	collector := colly.NewCollector()

	collector.OnHTML("tbody", func(element *colly.HTMLElement) {
		element.ForEach("tr", func(_ int, element *colly.HTMLElement) {
			Titles = append(Titles, element.ChildText(".text-bold.text-large1"))
		})
	})

	collector.Visit("https://imdb-api.com/most-popular-movies")

	for index := range Titles {
		filter := bson.M{"title": bson.M{"$eq": Titles[index]}}
		update := bson.M{"$set": bson.M{"popular": true}}
		collection.UpdateOne(ctx, filter, update)
	}
}


func GetPopularTvShows() {
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Movies")
	var Titles []string
	collector := colly.NewCollector()

	collector.OnHTML("tbody", func(element *colly.HTMLElement) {
		element.ForEach("tr", func(_ int, element *colly.HTMLElement) {
			Titles = append(Titles, element.ChildText(".text-bold.text-large1"))
		})
	})

	collector.Visit("https://imdb-api.com/most-popular-tvs")

	for index := range Titles {
		filter := bson.M{"title": bson.M{"$eq": Titles[index]}}
		update := bson.M{"$set": bson.M{"popular": true}}
		collection.UpdateOne(ctx, filter, update)
	}
}