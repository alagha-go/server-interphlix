package movies

import (
	"encoding/json"
	"interphlix/lib/movies"
	"interphlix/lib/variables"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/// http handler to handle update request to add server to movie.Server
func AddServer(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	if req.Method != "PUT" &&  req.Method != "UPDATE" && req.Method != "PATCH"{
		res.WriteHeader(http.StatusNotFound)
		return
	}
	ID, err := primitive.ObjectIDFromHex(mux.Vars(req)["id"])
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write(variables.JsonMarshal(variables.Error{Error: "invalid id"}))
		return
	}
	var Server movies.Server
	err = json.NewDecoder(req.Body).Decode(&Server)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write(variables.JsonMarshal(variables.Error{Error: "invalid json body"}))
		return
	}
	data, status := movies.AddServer(ID, Server)
	res.WriteHeader(status)
	res.Write(data)
}


/// http handler to handle get request to add url to movies urls
func AddUrl(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	if req.Method != "GET"{
		res.WriteHeader(http.StatusNotFound)
		return
	}
	params := mux.Vars(req)
	ID, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write(variables.JsonMarshal(variables.Error{Error: "invalid id"}))
		return
	}
	url := params["urls"]
	urls := strings.Split(url, ",")
	data, status := movies.AddUrl(ID, urls...)
	res.WriteHeader(status)
	res.Write(data)
}