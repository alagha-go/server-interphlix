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

func AddEpisodeUrl(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	params := mux.Vars(req)
	ID, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write(variables.JsonMarshal(variables.Error{Error: "invalid id"}))
		return
	}
	url := params["url"]
	seasonCode := params["seasoncode"]
	EpisodeCode := params["episodecode"]
	urls := strings.Split(url, ",")
	data, status := movies.AddEpioseUrl(ID, seasonCode, EpisodeCode, urls...)
	res.WriteHeader(status)
	res.Write(data)
}


func SetEpisodeServer(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	params := mux.Vars(req)
	ID, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write(variables.JsonMarshal(variables.Error{Error: "invalid id"}))
		return
	}
	data, status := movies.SetEpisodeServer(ID, params["seasoncode"], params["episodecode"], params["servername"])
	res.WriteHeader(status)
	res.Write(data)
}

func AddEpisodeServer(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
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
		res.Write(variables.JsonMarshal(variables.Error{Error: "invalid json body"}))
		return
	}
	data, status := movies.AddEpisodeServer(ID, params["seasoncode"], params["episodecode"], Server)
	res.WriteHeader(status)
	res.Write(data)
}