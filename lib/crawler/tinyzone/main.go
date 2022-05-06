package tinyzone

import (
	"interphlix/lib/crawler/tinyzone/movies"
	"interphlix/lib/crawler/tinyzone/tvshows"
)


func Main() {
	go movies.Main()
	go tvshows.Main()
}