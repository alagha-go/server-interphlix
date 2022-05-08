package movies


var (
	Movies []Movie
	Genres []Genre
)


func Main() {
	LoadMovies()
	LoadGenres()
	go ListenMovies()
}