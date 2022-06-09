package movies

import (
	"interphlix/lib/handler/accounts"
	"interphlix/lib/movies"
	"interphlix/lib/variables"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)


func GetMoviesByCast(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	valid := accounts.ValidateRequest(req)
	if !valid {
		res.WriteHeader(http.StatusUnauthorized)
		res.Write(variables.JsonMarshal(variables.Error{Error: "unauthorized"}))
		return
	}
	cast := mux.Vars(req)["cast"]
	round, err := strconv.Atoi(req.URL.Query().Get("round"))
	if err != nil {
		round = 0
	}
	data, status := movies.GetMoviesByCast(cast, round)
	res.WriteHeader(status)
	res.Write(data)
}