package movies

import (
	"interphlix/lib/handler/accounts"
	"interphlix/lib/movies"
	"interphlix/lib/movies/casts"
	"interphlix/lib/variables"
	"net/http"
	"strconv"
)


func SearchMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	err, status := accounts.ValidateRequest(req, "user")
	if err != nil {
		res.WriteHeader(status)
		res.Write(variables.JsonMarshal(variables.Error{Error: err.Error()}))
		return
	}
	Type := req.URL.Query().Get("type")
	query := req.URL.Query().Get("query")
	Genre := req.URL.Query().Get("genre")
	round, err := strconv.Atoi(req.URL.Query().Get("round"))
	if err != nil {
		round = 0
	}
	
	data, status := movies.SearchMovies(query, Type, Genre, round)
	res.WriteHeader(status)
	res.Write(data)
}


func SearchCast(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	err, status := accounts.ValidateRequest(req, "user")
	if err != nil {
		res.WriteHeader(status)
		res.Write(variables.JsonMarshal(variables.Error{Error: err.Error()}))
		return
	}
	query := req.URL.Query().Get("query")
	round, err := strconv.Atoi(req.URL.Query().Get("round"))
	if err != nil {
		round = 0
	}
	data, status := casts.SearchCasts(query, round)
	res.WriteHeader(status)
	res.Write(data)
}