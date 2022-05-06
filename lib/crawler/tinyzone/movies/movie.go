package movies

import (
	"interphlix/lib/types"
	"log"
	"strings"

	"github.com/gocolly/colly"
)


func CollectMovie(Movie *types.Movie) {
    if len(Movie.Servers) == 0 {
        CollectMovieContent(Movie)
    }
}

func CollectMovieContent(Movie *types.Movie) {
    collector := colly.NewCollector()

    collector.OnHTML(".description", func(element *colly.HTMLElement){
        Movie.Description = element.Text
        Movie.Description = strings.ReplaceAll(Movie.Description, "\n", "")
        Movie.Description = strings.ReplaceAll(Movie.Description, "  ", "")
        Movie.Description = strings.TrimPrefix(Movie.Description, " ")
        Movie.Description = strings.TrimSuffix(Movie.Description, " ")
    })

    SetServers(Movie)
    AddServer(Movie)

    collector.OnHTML(".elements", func(element *colly.HTMLElement) {
        SetElements(element, Movie)
    })

    SetID(Movie)
    collector.Visit(Movie.PageUrl)
}


func SetElements(element *colly.HTMLElement, Movie *types.Movie) {
    functions := []func(element *colly.HTMLElement, movie *types.Movie){}
    functions = append(functions,  SetReleased)
    functions = append(functions,  SetGenre)
    functions = append(functions,  SetCasts)
    functions = append(functions,  SetDuration)
    functions = append(functions,  SetCountries)
    functions = append(functions,  SetProducers)
    element.ForEach(".row-line", func(index int, element *colly.HTMLElement){
        functions[index](element, Movie)
    })
}


func SetServers(Movie *types.Movie) {
    collector := colly.NewCollector()

    url := "https://tinyzonetv.to/ajax/movie/episodes/"+ Movie.Code

    collector.OnHTML(".nav", func(element *colly.HTMLElement) {
        var servers []types.Server
        element.ForEach(".nav-item", func(index int, element *colly.HTMLElement) {
            var server types.Server
            server.WatchID = element.ChildAttr("a", "data-linkid")
            server.Name = element.ChildAttr("a", "title")
            server.Name = strings.ReplaceAll(server.Name, "Server ", "")
            servers = append(servers, server)
        })
        Movie.Servers = servers
    })
    collector.Visit(url)
}


func SetID(Movie *types.Movie) {
	for index, server := range Movie.Servers {
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
			Movie.Servers[index].ID = strings.ReplaceAll(res.Link, "https://streamlare.com/e/", "")
			Movie.Servers[index].Url = "https://streamlare.com/v/" + Movie.Servers[index].ID
            Movie.Server = &server
            Available++
		}else if server.Name == "Vidcloud" || server.Name == "UpCloud"{
			Movie.Servers[index].ID = strings.ReplaceAll(res.Link, "https://rabbitstream.net/embed-4/", "")
			Movie.Servers[index].ID = strings.ReplaceAll(Movie.Servers[index].ID, "?z=", "")
			Movie.Servers[index].Url = "https://rabbitstream.net/embed/m-download/" + Movie.Servers[index].ID
		}else {
			Movie.Servers[index].Url = res.Link
		}
	}
}

func AddServer(Movie *types.Movie) {
    for _, server := range Movie.Servers {
        if server.Name == "Vidcloud" || server.Name == "UpCloud" {
            collector := colly.NewCollector()

			collector.OnHTML("#user_menu", func(element *colly.HTMLElement) {

            })
			collector.Visit(server.Url)
        }
    }
}

func AddServers(element *colly.HTMLElement ,Movie *types.Movie) {
    element.ForEach(".dl-site", func(_ int, element *colly.HTMLElement) {
		var exist bool = false
		var server types.Server
		server.Name = element.ChildText(".site-name")
		server.Url = element.ChildAttr("a", "href")
		for index, serve := range Movie.Servers {
			if serve.Name == server.Name {
				Movie.Servers[index].Url = server.Url
				exist = true
			}
		}
		if !exist {
			Movie.Servers = append(Movie.Servers, server)
		}
	})
}