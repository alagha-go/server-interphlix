package movies

import (
	"encoding/json"
	"interphlix/lib/handler/accounts"
	"interphlix/lib/movies"
	"interphlix/lib/variables"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func AddEpisode(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	valid := accounts.ValidateRequest(req)
	if !valid {
		res.WriteHeader(http.StatusUnauthorized)
		res.Write(variables.JsonMarshal(variables.Error{Error: "unauthorized"}))
		return
	}
	params := mux.Vars(req)
	MovieID, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write(variables.JsonMarshal(variables.Error{Error: "invalid movie id"}))
	}
	SeasonID, err := primitive.ObjectIDFromHex(params["seasonid"])
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write(variables.JsonMarshal(variables.Error{Error: "invalid season id"}))
	}
	var Episode movies.Episode
	json.NewDecoder(req.Body).Decode(&Episode)
	data, status := movies.AddEpisode(MovieID, SeasonID, Episode)
	res.WriteHeader(status)
	res.Write(data)
}