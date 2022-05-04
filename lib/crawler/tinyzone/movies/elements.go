package movies

import (
	"interphlix/lib/types"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)



func SetReleased(element *colly.HTMLElement, Movie *types.Movie) {
    layout := "2006-01-02"
    released := element.Text
    released = strings.ReplaceAll(released, "Released: ", "")
    released = strings.ReplaceAll(released, "  ", "")
    released = strings.ReplaceAll(released, "\n", "")
    released = strings.TrimPrefix(released, " ")
    released = strings.TrimSuffix(released, " ")
    Released, _ := time.Parse(layout, released)
    Movie.Released = &Released
}

func SetGenre(element *colly.HTMLElement, Movie *types.Movie) {
    element.ForEach("a", func(index int, element *colly.HTMLElement){
        Movie.Genres = append(Movie.Genres, element.Text)
    })
}

func SetCasts(element *colly.HTMLElement, Movie *types.Movie) {
    element.ForEach("a", func(index int, element *colly.HTMLElement){
        Movie.Casts = append(Movie.Casts, element.Text)
    })
}

func SetDuration(element *colly.HTMLElement, Movie *types.Movie) {
    duration := element.Text
    duration = strings.ReplaceAll(duration, "Duration: ", "")
    duration = strings.ReplaceAll(duration, "  ", "")
    duration = strings.ReplaceAll(duration, "\n", "")
    duration = strings.ReplaceAll(duration, "min", "")
    duration = strings.TrimPrefix(duration, " ")
    duration = strings.TrimSuffix(duration, " ")
    if strings.Contains(duration, "N/A") {
        Movie.Duration = 0
    }else {
        minutes, _ := strconv.Atoi(duration)
        Movie.Duration = minutes
    }
}

func SetCountries(element *colly.HTMLElement, Movie *types.Movie) {
    element.ForEach("a", func(index int, element *colly.HTMLElement){
        Movie.Countries = append(Movie.Countries, element.Text)
    })
}

func SetProducers(element *colly.HTMLElement, Movie *types.Movie) {
    production := element.Text
    production = strings.ReplaceAll(production, "Production: ", "")
    production = strings.ReplaceAll(production, "  ", "")
    production = strings.ReplaceAll(production, "\n", "")
    production = strings.TrimPrefix(production, " ")
    production = strings.TrimSuffix(production, " ")
    if production == "N/A" {
        return
    }
    Movie.Producers = strings.Split(production, ",")
}