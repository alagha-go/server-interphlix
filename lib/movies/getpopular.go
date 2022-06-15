package movies

import (
	"interphlix/lib/variables"
	"net/http"
)


func GetPoPularMovies(seed int64, round int) ([]byte, int) {
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

	if start > len(Movies) {
		return variables.JsonMarshal([]Movie{}), http.StatusOK
	}else if end > len(Movies) {
		return variables.JsonMarshal(Movies[start:]), http.StatusOK
	}

	return variables.JsonMarshal(Movies[start:end]), http.StatusOK
}


func GetPoPularTvShows(seed int64, round int) ([]byte, int) {
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

	if start > len(Movies) {
		return variables.JsonMarshal([]Movie{}), http.StatusOK
	}else if end > len(Movies) {
		return variables.JsonMarshal(Movies[start:]), http.StatusOK
	}

	return variables.JsonMarshal(Movies[start:end]), http.StatusOK
}