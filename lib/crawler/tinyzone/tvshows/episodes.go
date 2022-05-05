package tvshows

import (
	"interphlix/lib/types"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)


func GetEpisodes(Season *types.Season) {
	collector := colly.NewCollector()
	url := "https://tinyzonetv.to/ajax/v2/season/episodes/" + Season.Code

	collector.OnHTML(".nav", func(element *colly.HTMLElement) {
		CollectAllEpisodes(element, Season)
	})

	collector.Visit(url)
}

func CollectAllEpisodes(element *colly.HTMLElement, Season *types.Season) {
	element.ForEach(".nav-item", func(_ int, element *colly.HTMLElement) {
		var Episode types.Episode
		Episode.Name = element.ChildText("a")
		index := strings.Index(Episode.Name, "s")
		end := strings.Index(Episode.Name, "\n")
		Episode.Index, _ = strconv.Atoi(Episode.Name[index:end])
		Episode.Name = strings.ReplaceAll(Episode.Name, `Eps 9
                    : `, "")
		Episode.Code = element.ChildAttr("a", "data-id")
		GetAllServers(&Episode)
		Season.Episodes = append(Season.Episodes, Episode)
	})
}

func GetAllServers(Episode *types.Episode) {
	collector := colly.NewCollector()
	url := "https://tinyzonetv.to/ajax/v2/episode/servers/" + Episode.Code

	collector.OnHTML(".nav", func(element *colly.HTMLElement) {
		CollectServers(element, Episode)
	})

	collector.Visit(url)
}

func CollectServers(element *colly.HTMLElement, Episode *types.Episode) {
	element.ForEach(".nav-item", func(_ int, element *colly.HTMLElement) {
		var Server types.Server
		Server.WatchID = element.ChildAttr("a", "data-id")
		Server.Name = element.ChildAttr("a", "title")
		Server.Name = strings.ReplaceAll(Server.Name, "Server ", "")
		AddServer(Episode)
		SetID(Episode)
	})
}