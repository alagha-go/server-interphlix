package movies

import (
	"context"
	"interphlix/lib/movies/casts"
	"interphlix/lib/movies/genres"
	"interphlix/lib/movies/types"
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