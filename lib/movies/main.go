package movies

import (
	"interphlix/lib/movies/casts"
	"interphlix/lib/movies/genres"
	"interphlix/lib/movies/types"
)


var (
	Movies []Movie
)


func Main() {
	LoadMovies()
	genres.LoadGenres()
	types.Main()
	genres.Main()
	casts.Main()
	Listener()
}


func LoadMoviesFromDB() {

}