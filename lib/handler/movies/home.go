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
	err, status := accounts.ValidateRequest(req, "user")
	if err != nil {
		res.WriteHeader(status)
		res.Write(variables.JsonMarshal(variables.Error{Error: err.Error()}))
		return
	}
	if req.Method != "GET" {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	cookie, _ := req.Cookie("token")
	account, err := accounts.GetAccount(cookie.Value)
	data, status := movies.GetHome(account.ID)
	res.WriteHeader(status)
	res.Write(data)
}


func GetMoviesByGenre(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	err, status := accounts.ValidateRequest(req, "user")
	if err != nil {
		res.WriteHeader(status)
		res.Write(variables.JsonMarshal(variables.Error{Error: err.Error()}))
		return
	}
	genre := mux.Vars(req)["genre"]
	round, err := strconv.Atoi(req.URL.Query().Get("round"))
	if err != nil {
		round = 0
	}
	seed, err := strconv.ParseInt(req.URL.Query().Get("seed"), 10, 64)
	if err != nil {
		seed = 0
	}
	data, status := movies.GetMoviesByGenre(genre, round, seed)
	res.WriteHeader(status)
	res.Write(data)
}