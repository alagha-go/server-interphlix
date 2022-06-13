package movies

import (
	"encoding/json"
	"interphlix/lib/handler/accounts"
	"interphlix/lib/movies"
	"interphlix/lib/variables"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetMoviesByTypeAndGenre(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	err, status := accounts.ValidateRequest(req, "user")
	if err != nil {
		res.WriteHeader(status)
		res.Write(variables.JsonMarshal(variables.Error{Error: err.Error()}))
		return
	}
	Type := mux.Vars(req)["type"]
	Genre := mux.Vars(req)["genre"]

	round, err := strconv.Atoi(req.URL.Query().Get("round"))
	if err != nil {
		round = 0
	}

	data, status := movies.GetMoviesByGenreAndType(Type, Genre, round)

	res.WriteHeader(status)
	res.Write(data)
}


func GetMovieByID(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	err, status := accounts.ValidateRequest(req, "user")
	if err != nil {
		res.WriteHeader(status)
		res.Write(variables.JsonMarshal(variables.Error{Error: err.Error()}))
		return
	}
	ID, err := primitive.ObjectIDFromHex(mux.Vars(req)["id"])
	if err != nil {
		variables.JsonMarshal(variables.Error{Error: "invalid id"})
	}
	Movie, err := movies.GetMovieByID(ID)
	if err != nil {
		variables.JsonMarshal(variables.Error{Error: err.Error()})
	}
	json.NewEncoder(res).Encode(Movie)
}