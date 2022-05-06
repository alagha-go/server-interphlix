package tvshows

import (
	"interphlix/lib/types"
	"io/ioutil"
	"strings"

	"github.com/gocolly/colly"
)


var (
	TvShows []types.Movie
	Position int
	Available int
)

func CollectAllPages(pages int) {
	for index:=1; index<pages+1; index++ {
		CollectPage(index)
	}
	SaveTvShows()
	types.PrintGreen(len(TvShows))
	types.PrintCyan("done collecting tvshows from all pages")

	for index := range TvShows{
		CollectTvShow(&TvShows[index])
		Position = index+1
		if TvShows[index].Seasons[0].Episodes[0].Available {
			Available++
		}
		SaveTvShows()
	}
}

func CollectPage(number int) {
	url := "https://tinyzonetv.to/tv-show?page=" + string(rune(number))
	collector := colly.NewCollector()

	collector.OnHTML(".film_list-wrap", CollectTvShows)

	collector.Visit(url)
}

func CollectTvShows(element *colly.HTMLElement) {
	element.ForEach(".flw-item", func(_ int, element *colly.HTMLElement) {
		var TvShow types.Movie
        TvShow.Title = element.ChildAttr("a", "title")
        TvShow.ImageUrl = element.ChildAttr("img", "data-src")
        TvShow.PageUrl = "https://tinyzonetv.to" + element.ChildAttr("a", "href")
		index := strings.Index(TvShow.PageUrl, "free-")
    	TvShow.Code = TvShow.PageUrl[index+5:]
		GetSeasons(&TvShow)
		GetTvShowUrls(&TvShow)
		TvShows = append(TvShows, TvShow)
		SaveTvShows()
	})
}

func SaveTvShows() {
	data := types.JsonMarshal(TvShows)
	ioutil.WriteFile("./DB/tvshows.json", data, 0755)
}