package movies

import (
	"interphlix/lib/movies/genres"
	"interphlix/lib/variables"
	"net/http"

	"github.com/gorilla/mux"
)


func GetAllGenres(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	if req.Method != "GET" {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	res.WriteHeader(200)
	res.Write(variables.JsonMarshal(genres.Genres))
}


func GetGenreByType(res http.ResponseWriter, req *http.Request) {
	var Genres []genres.Genre
	res.Header().Set("content-type", "application/json")
	if req.Method != "GET" {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	Type := mux.Vars(req)["type"]
	for index := range genres.Genres {
		for tindex := range genres.Genres[index].Types {
			if genres.Genres[index].Types[tindex] == Type{
				Genres = append(Genres, genres.Genres[index])
			}
		}
	}
	res.WriteHeader(200)
	res.Write(variables.JsonMarshal(Genres))
}