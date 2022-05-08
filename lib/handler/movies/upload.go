package movies

import (
	"encoding/json"
	"interphlix/lib/movies"
	"interphlix/lib/variables"
	"net/http"
)


func UploadMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	if req.Method != "POST" {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	var Movie movies.Movie
	err := json.NewDecoder(req.Body).Decode(&Movie)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write(variables.JsonMarshal(variables.Error{Error: "invalid json"}))
		return
	}
	data, status := movies.UploadOneMovie(Movie)
	res.WriteHeader(status)
	res.Write(data)
}