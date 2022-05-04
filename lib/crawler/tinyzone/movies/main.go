package movies

import (
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func Main() {
	CollectAllPages(GetNumberOfPages())
}


func GetNumberOfPages() int {
	var err error
	var numberofPages int
	collector := colly.NewCollector()

	collector.OnHTML(".pagination.pagination-lg.justify-content-center", func(element *colly.HTMLElement) {
		element.ForEach(".page-item", func(_ int, element *colly.HTMLElement) {
			title := element.ChildAttr("a", "title")
			href := element.ChildAttr("a", "href")
			if title == "Last" {
				href = strings.ReplaceAll(href, "/movie?page=", "")
				numberofPages, err = strconv.Atoi(href)
				HanleError(err)
			}
		})
	})

	collector.Visit("https://tinyzonetv.to/movie")

	return numberofPages
}

func HanleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}