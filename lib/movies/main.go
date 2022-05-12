package movies

import "interphlix/lib/movies/genres"


var (
	Movies []Movie
)


func Main() {
	LoadMovies()
	genres.LoadGenres()
}


func LoadMoviesFromDB() {

}