package movies

import (
	"interphlix/lib/handler/accounts"
	"interphlix/lib/movies/genres"
	"interphlix/lib/variables"
	"net/http"

	"github.com/gorilla/mux"
)


func GetAllGenres(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	err, status := accounts.ValidateRequest(req, "user")
	if err != nil {
		res.WriteHeader(status)
		res.Write(variables.JsonMarshal(variables.Error{Error: err.Error()}))
		return
	}
	res.WriteHeader(200)
	res.Write(variables.JsonMarshal(genres.GetAllGenres()))
}


func GetGenreByType(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	err, status := accounts.ValidateRequest(req, "user")
	if err != nil {
		res.WriteHeader(status)
		res.Write(variables.JsonMarshal(variables.Error{Error: err.Error()}))
		return
	}
	Type := mux.Vars(req)["type"]
	res.WriteHeader(200)
	res.Write(variables.JsonMarshal(genres.GetGenresByType(Type)))
}