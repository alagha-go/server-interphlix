package genres

var (
	Genres []Genre
)

func Main() {
	LoadGenres()
	Listener()
}