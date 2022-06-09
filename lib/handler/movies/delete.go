package movies

import (
	"encoding/json"
	"interphlix/lib/handler/accounts"
	"interphlix/lib/movies"
	"interphlix/lib/variables"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func DeleteUrls(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	err, status := accounts.ValidateRequest(req, "user")
	if err != nil {
		res.WriteHeader(status)
		res.Write(variables.JsonMarshal(variables.Error{Error: err.Error()}))
		return
	}
	params := mux.Vars(req)
	ID, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write(variables.JsonMarshal(variables.Error{Error: "invalid id"}))
		return
	}
	urls := strings.Split(params["urls"], ",")
	data, status := movies.DeleteUrls(ID, params["seasoncode"], params["episodecode"], urls...)
	res.WriteHeader(status)
	res.Write(data)
}

func DeleteServer(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	err, status := accounts.ValidateRequest(req, "user")
	if err != nil {
		res.WriteHeader(status)
		res.Write(variables.JsonMarshal(variables.Error{Error: err.Error()}))
		return
	}
	params := mux.Vars(req)
	ID, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write(variables.JsonMarshal(variables.Error{Error: "invalid id"}))
		return
	}
	var Server movies.Server
	err = json.NewDecoder(req.Body).Decode(&Server)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write(variables.JsonMarshal(variables.Error{Error: "invalid json"}))
		return
	}
	data, status := movies.DeleteServer(ID, params["seasoncode"], params["episodecode"], Server)
	res.WriteHeader(status)
	res.Write(data)
}