package movies

import (
	"interphlix/lib/handler/accounts"
	"interphlix/lib/movies"
	"interphlix/lib/variables"
	"net/http"
	"strconv"

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
	round, err := strconv.Atoi(req.URL.Query().Get("round"))
	if err != nil {
		round = 0
	}
	data, status := movies.GetMovies(round)
	res.WriteHeader(status)
	res.Write(data)
}


func GetMoviesByGenre(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	valid := accounts.ValidateRequest(req)
	if !valid {
		res.WriteHeader(http.StatusUnauthorized)
		res.Write(variables.JsonMarshal(variables.Error{Error: "unauthorized"}))
		return
	}
	genre := mux.Vars(req)["genre"]
	round, err := strconv.Atoi(req.URL.Query().Get("round"))
	if err != nil {
		round = 0
	}
	data, status := movies.GetMoviesByGenre(genre, round)
	res.WriteHeader(status)
	res.Write(data)
}