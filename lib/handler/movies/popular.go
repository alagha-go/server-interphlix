package movies

import (
	"interphlix/lib/handler/accounts"
	"interphlix/lib/movies"
	"interphlix/lib/variables"
	"net/http"
	"strconv"
	"time"
)


func GetPopularMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	err, status := accounts.ValidateRequest(req, "user")
	if err != nil {
		res.WriteHeader(status)
		res.Write(variables.JsonMarshal(variables.Error{Error: err.Error()}))
		return
	}
	round, err := strconv.Atoi(req.URL.Query().Get("round"))
	if err != nil {
		round = 1
	}
	seed, err := strconv.ParseInt(req.URL.Query().Get("seed"), 10, 64)
	if err != nil {
		seed = time.Now().UnixNano()
		round = 0
	}
	data, status := movies.GetPoPularMovies(seed, round)
	res.WriteHeader(status)
	res.Write(data)
}


func GetPopularTvShows(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	err, status := accounts.ValidateRequest(req, "user")
	if err != nil {
		res.WriteHeader(status)
		res.Write(variables.JsonMarshal(variables.Error{Error: err.Error()}))
		return
	}
	round, err := strconv.Atoi(req.URL.Query().Get("round"))
	if err != nil {
		round = 1
	}
	seed, err := strconv.ParseInt(req.URL.Query().Get("seed"), 10, 64)
	if err != nil {
		seed = time.Now().UnixNano()
		round = 0
	}
	data, status := movies.GetPoPularTvShows(seed, round)
	res.WriteHeader(status)
	res.Write(data)
}


func GetFeatured(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	err, status := accounts.ValidateRequest(req, "user")
	if err != nil {
		res.WriteHeader(status)
		res.Write(variables.JsonMarshal(variables.Error{Error: err.Error()}))
		return
	}
	round, err := strconv.Atoi(req.URL.Query().Get("round"))
	if err != nil {
		round = 1
	}
	seed, err := strconv.ParseInt(req.URL.Query().Get("seed"), 10, 64)
	if err != nil {
		seed = time.Now().UnixNano()
		round = 0
	}
	data, status := movies.GetFeatured(seed, round)
	res.WriteHeader(status)
	res.Write(data)
}


func GetTrending(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	err, status := accounts.ValidateRequest(req, "user")
	if err != nil {
		res.WriteHeader(status)
		res.Write(variables.JsonMarshal(variables.Error{Error: err.Error()}))
		return
	}
	round, err := strconv.Atoi(req.URL.Query().Get("round"))
	if err != nil {
		round = 0
	}
	seed, err := strconv.ParseInt(req.URL.Query().Get("seed"), 10, 64)
	if err != nil {
		seed = time.Now().UnixNano()
		round = 0
	}
	data, status := movies.GetTrendingMoviesApi(seed, round)
	res.WriteHeader(status)
	res.Write(data)
}