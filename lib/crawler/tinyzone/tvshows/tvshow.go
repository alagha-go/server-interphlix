package tvshows

import (
	"interphlix/lib/types"
	"strings"

	"github.com/gocolly/colly"
)


func CollectTvShow(TvShow *types.Movie) {
    if len(TvShow.Servers) == 0 {
        CollectTvShowContent(TvShow)
    }
}

func CollectTvShowContent(TvShow *types.Movie) {
    collector := colly.NewCollector()

    collector.OnHTML(".description", func(element *colly.HTMLElement){
        TvShow.Description = element.Text
        TvShow.Description = strings.ReplaceAll(TvShow.Description, "\n", "")
        TvShow.Description = strings.ReplaceAll(TvShow.Description, "  ", "")
        TvShow.Description = strings.TrimPrefix(TvShow.Description, " ")
        TvShow.Description = strings.TrimSuffix(TvShow.Description, " ")
    })

    collector.OnHTML(".elements", func(element *colly.HTMLElement) {
        SetElements(element, TvShow)
    })
    collector.Visit(TvShow.PageUrl)
}


func SetElements(element *colly.HTMLElement, TvShow *types.Movie) {
    functions := []func(element *colly.HTMLElement, TvShow *types.Movie){}
    functions = append(functions,  SetReleased)
    functions = append(functions,  SetGenre)
    functions = append(functions,  SetCasts)
    functions = append(functions,  SetDuration)
    functions = append(functions,  SetCountries)
    functions = append(functions,  SetProducers)
    element.ForEach(".row-line", func(index int, element *colly.HTMLElement){
        functions[index](element, TvShow)
    })
}