package movies



var (
	Titles []string
)


func CollectMovies() {
	GetPopularMovies()
	GetPopularTvShows()
}