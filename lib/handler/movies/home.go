package movies

import (
	"interphlix/lib/handler/accounts"
	"interphlix/lib/movies"
	"interphlix/lib/variables"
	"net/http"

	"github.com/gorilla/mux"
)


func GetHomeMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	valid := accounts.ValidateRequest(req)
	if !valid {
		res.WriteHeader(http.StatusUnauthorized)
		res.Write(variables.JsonMarshal(variables.Error{Error: "unauthorized"}))
		return
	}
	if req.Method != "GET" {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	data, status := movies.GetMovies()
	res.WriteHeader(status)
	res.Write(data)
}


func GetMoviesByGenre(res http.ResponseWriter, req *http.Request) {
	var Movies []movies.Movie
	res.Header().Set("content-type", "application/json")
	valid := accounts.ValidateRequest(req)
	if !valid {
		res.WriteHeader(http.StatusUnauthorized)
		res.Write(variables.JsonMarshal(variables.Error{Error: "unauthorized"}))
		return
	}
	genre := mux.Vars(req)["genre"]
	for _, Movie := range movies.Movies {
		for _, Genre := range Movie.Genres {
			if genre == Genre {
				Movies = append(Movies, Movie)
			}
		}
	}
	res.WriteHeader(200)
	res.Write(variables.JsonMarshal(Movies))
}