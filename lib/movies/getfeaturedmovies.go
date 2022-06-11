package movies

import (
	"context"
	"fmt"
	"interphlix/lib/variables"

	"github.com/gocolly/colly"
	"go.mongodb.org/mongo-driver/bson"
)


var (
	Titles []string
	FeaturedMovies []Movie
)


func CollectFeaturesTitles() {
	start := 1
	for start < 10000{
		CollectData(start)
		start += 50
	}
	SetFeatured()
}

func CollectData(start int) {
	collector := colly.NewCollector()

	collector.OnHTML(".lister-list", func(element *colly.HTMLElement) {
		element.ForEach(".lister-item.mode-advanced", func(_ int, element *colly.HTMLElement) {
			element.ForEach("h3", func(_ int, element *colly.HTMLElement) {
				Titles = append(Titles, element.ChildText("a"))
			})
		})
	})

	collector.Visit(fmt.Sprintf("https://www.imdb.com/search/title/?title_type=feature&year=2010,2022&start=%d", start))
}


func SetFeatured() {
	ctx := context.Background()
	collection := variables.Client1.Database("Interphlix").Collection("Movies")
	for index := range Titles {
		var Movie Movie
		err := collection.FindOne(ctx, bson.M{"title": Titles[index]}).Decode(&Movie)
		if err == nil {
			FeaturedMovies = append(FeaturedMovies, Movie)
		}
	}
}