package movies

import (
	"encoding/json"
	"interphlix/lib/movies"
	"interphlix/lib/variables"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func SetServer(res http.ResponseWriter, req *http.Request) {
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
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write(variables.JsonMarshal(variables.Error{Error: "invalid json body"}))
		return
	}
	err = json.NewDecoder(req.Body).Decode(&Server)
	data, status := movies.SetServer(ID, Server)
	res.WriteHeader(status)
	res.Write(data)
}

func AddUrl(res http.ResponseWriter, req *http.Request) {
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
	url := params["url"]
	data, status := movies.AddUrl(ID, url)
	res.WriteHeader(status)
	res.Write(data)
}