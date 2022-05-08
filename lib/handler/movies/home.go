package movies

import (
	"interphlix/lib/movies"
	"net/http"
)


func GetHomeMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	if req.Method != "GET" {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	data, status := movies.GetMovies()
	res.WriteHeader(status)
	res.Write(data)
}