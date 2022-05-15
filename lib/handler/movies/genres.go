package movies

import (
	"interphlix/lib/movies/genres"
	"interphlix/lib/variables"
	"net/http"
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