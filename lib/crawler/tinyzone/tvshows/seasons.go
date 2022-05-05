package tvshows

import (
	"interphlix/lib/types"

	"github.com/gocolly/colly"
)


func GetSeasons(TvShow *types.Movie) {
	collector := colly.NewCollector()
	url := "https://tinyzonetv.to/ajax/v2/tv/seasons/" + TvShow.Code

	collector.OnHTML(".dropdown-menu.dropdown-menu-new", func(element *colly.HTMLElement) {
		CollectAllSeasons(element, TvShow)
	})

	collector.Visit(url)
}

func CollectAllSeasons(element *colly.HTMLElement, TvShow *types.Movie) {
	element.ForEach("a", func(index int, element *colly.HTMLElement) {
		var Season types.Season
		Season.Index = index
		Season.Code = element.Attr("data-id")
		Season.Name = element.Text
		GetEpisodes(&Season)
		TvShow.Seasons = append(TvShow.Seasons, Season)
	})
}