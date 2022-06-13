package movies

import (
	"context"
	"fmt"
	"interphlix/lib/variables"

	"github.com/gocolly/colly"
	"go.mongodb.org/mongo-driver/bson"
)


var (
	Names []string
)


func CollectFeatured() {
	index := 1;
	for len(Names) < 1000 {
		url := fmt.Sprintf("https://www.imdb.com/search/title/?title_type=feature&sort=user_rating,desc&start=%d", index)
		CollectPage(url)
		index += 50
	}
	SetFeatured()
}

func CollectPage(url string) {
	collector := colly.NewCollector()

	collector.OnHTML(".lister-list", func(element *colly.HTMLElement) {
		element.ForEach("h3", func(_ int, element *colly.HTMLElement) {
			Name := element.ChildText("a")
			Names = append(Names, Name)
		})
	})

	collector.Visit(url)
}


func SetFeatured() {
	ctx := context.Background()
	collection := variables.Client.Database("Interphlix").Collection("Movies")

	for index := range Names {
		movie := Movie{Title: Names[index]}
		exists, Movie := movie.ExistByTitle()
		if exists && !Movie.Featured {
			filter := bson.M{"_id": bson.M{"$eq": Movie.ID}}
			update := bson.M{"$set": bson.M{"featured": true}}
			collection.UpdateOne(ctx, filter, update)
		}
	}
}



func (movie *Movie) ExistByTitle() (bool, Movie) {
	var Movie Movie
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Movies")

	err := collection.FindOne(ctx, bson.M{"title": movie.Title}).Decode(&Movie)
	return err == nil, Movie
}