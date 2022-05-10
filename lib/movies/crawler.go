package movies

import (
	"time"

	"github.com/gocolly/colly"
)


func (Movie *Movie) AddServers() {
	if len(Movie.Servers) > 0 {
		for _, Server := range Movie.Servers {
			if Server.Name == "Vidcloud" {
						collector := colly.NewCollector()

						collector.OnHTML("#user_menu", func(element *colly.HTMLElement) {
							servers := GetServers(element)
							Movie.Servers = append(Movie.Servers, servers...)
						})
						collector.Visit(Server.Url)
					}
		}
	}else {
		for Sindex, Season := range Movie.Seasons {
			for eindex, Episode := range Season.Episodes {
				for _, Server := range Episode.Servers {
					if Server.Name == "Vidcloud" {
						collector := colly.NewCollector()

						collector.OnHTML("#user_menu", func(element *colly.HTMLElement) {
							servers := GetServers(element)
							Movie.Seasons[Sindex].Episodes[eindex].Servers = append(Movie.Seasons[Sindex].Episodes[eindex].Servers, servers...)
						})
						collector.Visit(Server.Url)
					}
				}
			}
		}
	}
}

func GetServers(element *colly.HTMLElement) []Server {
	time.Sleep(500*time.Millisecond)
	var Servers []Server
    element.ForEach(".dl-site", func(_ int, element *colly.HTMLElement) {
		var server Server
		server.Name = element.ChildText(".site-name")
		server.Url = element.ChildAttr("a", "href")
		Servers = append(Servers, server)
	})
	return Servers
}