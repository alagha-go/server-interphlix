package movies

import (
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
	Listener()
}


func LoadMoviesFromDB() {

}