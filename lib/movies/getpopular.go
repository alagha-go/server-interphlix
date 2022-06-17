package movies

import (
	"interphlix/lib/variables"
	"net/http"
)


func GetPoPularMovies(seed int64, round int) ([]byte, int) {
	var movies []Movie
	start := 0
	end := 20;
	if round != 0 {
		start = (round*20) + 1
		end = (round*20) + 21
	}
	Movies, err := GetPopularMovies()
	if err != nil {
		return variables.JsonMarshal(variables.Error{Error: err.Error()}), http.StatusInternalServerError
	}
	Movies = RandomMovies(seed, Movies)

	if start > len(Movies) {
		return variables.JsonMarshal([]Movie{}), http.StatusOK
	}else if end > len(Movies) {
		Movies = Movies[start:]
	}else {
		Movies = Movies[start:end]
	}

	for _, movie := range Movies {
		movies = append(movies, Movie{ID: movie.ID, Code: movie.Code, Title: movie.Title, Type: movie.Type, ImageUrl: movie.ImageUrl})
	}

	return variables.JsonMarshal(movies), http.StatusOK
}


func GetPoPularTvShows(seed int64, round int) ([]byte, int) {
	var movies []Movie
	start := 0
	end := 20;
	if round != 0 {
		start = (round*20) + 1
		end = (round*20) + 21
	}
	Movies, err := GetPopularTvShows()
	if err != nil {
		return variables.JsonMarshal(variables.Error{Error: err.Error()}), http.StatusInternalServerError
	}
	Movies = RandomMovies(seed, Movies)

	if start > len(Movies) {
		return variables.JsonMarshal([]Movie{}), http.StatusOK
	}else if end > len(Movies) {
		Movies = Movies[start:]
	}else {
		Movies = Movies[start:end]
	}

	for _, movie := range Movies {
		movies = append(movies, Movie{ID: movie.ID, Code: movie.Code, Title: movie.Title, Type: movie.Type, ImageUrl: movie.ImageUrl})
	}

	return variables.JsonMarshal(movies), http.StatusOK
}


func GetFeatured(seed int64, round int) ([]byte, int) {
	var movies []Movie
	start := 0
	end := 20;
	if round != 0 {
		start = (round*20) + 1
		end = (round*20) + 21
	}
	Movies, err := GetFeaturedMovies(seed)
	if err != nil {
		return variables.JsonMarshal(variables.Error{Error: err.Error()}), http.StatusInternalServerError
	}

	Movies = RandomMovies(seed, Movies)

	if start > len(Movies) {
		return variables.JsonMarshal([]Movie{}), http.StatusOK
	}else if end > len(Movies) {
		Movies = Movies[start:]
	}else {
		Movies = Movies[start:end]
	}

	for _, movie := range Movies {
		movies = append(movies, Movie{ID: movie.ID, Code: movie.Code, Title: movie.Title, Type: movie.Type, ImageUrl: movie.ImageUrl})
	}

	return variables.JsonMarshal(movies), http.StatusOK
}