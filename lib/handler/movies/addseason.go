package movies

import (
	"encoding/json"
	"interphlix/lib/movies"
	"interphlix/lib/variables"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func AddSeason(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	if req.Method != "POST" {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	params := mux.Vars(req)
	ID, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write(variables.JsonMarshal(variables.Error{Error: "invalid movie id"}))
	}
	var Season movies.Season
	json.NewDecoder(req.Body).Decode(&Season)
	data, status := movies.AddSeason(ID, Season)
	res.WriteHeader(status)
	res.Write(data)
}