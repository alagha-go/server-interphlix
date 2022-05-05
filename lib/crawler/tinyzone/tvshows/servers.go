package tvshows

import (
	"interphlix/lib/types"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func SetID(Episode *types.Episode) {
	for index, server := range Episode.Servers {
		url := "https://tinyzonetv.to/ajax/get_link/"+ server.WatchID
		data, _, err := GetRequest(url, false)
		if err != nil {
			log.Panic(err)
		}
		res, err := UnmarshalLinkResponse(data)
		if err != nil {
			log.Panic(err)
		}
        if server.Name == "Streamlare" {
			Episode.Servers[index].ID = strings.ReplaceAll(res.Link, "https://streamlare.com/e/", "")
			Episode.Servers[index].Url = "https://streamlare.com/v/" + Episode.Servers[index].ID
            Episode.Server = server
            Available++
		}else if server.Name == "Vidcloud" || server.Name == "UpCloud"{
			Episode.Servers[index].ID = strings.ReplaceAll(res.Link, "https://rabbitstream.net/embed-4/", "")
			Episode.Servers[index].ID = strings.ReplaceAll(Episode.Servers[index].ID, "?z=", "")
			Episode.Servers[index].Url = "https://rabbitstream.net/embed/m-download/" + Episode.Servers[index].ID
		}else {
			Episode.Servers[index].Url = res.Link
		}
	}
}

func AddServer(Episode *types.Episode) {
    for _, server := range Episode.Servers {
        if server.Name == "Vidcloud" || server.Name == "UpCloud" {
            collector := colly.NewCollector()

			collector.OnHTML("#user_menu", func(element *colly.HTMLElement) {

            })
			collector.Visit(server.Url)
        }
    }
}

func AddServers(element *colly.HTMLElement ,Episode *types.Episode) {
    element.ForEach(".dl-site", func(_ int, element *colly.HTMLElement) {
		var exist bool = false
		var server types.Server
		server.Name = element.ChildText(".site-name")
		server.Url = element.ChildAttr("a", "href")
		for index, serve := range Episode.Servers {
			if serve.Name == server.Name {
				Episode.Servers[index].Url = server.Url
				exist = true
			}
		}
		if !exist {
			Episode.Servers = append(Episode.Servers, server)
		}
	})
}