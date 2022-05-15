package handler

import (
	"interphlix/lib/handler/movies"
	"interphlix/lib/servers"
	"interphlix/lib/variables"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)


var (
	Client *mongo.Client
	Router = mux.NewRouter()
)

func Main() {
	movies.Main()
	variables.LoadErrors()
	Router.HandleFunc("/", movies.GetHomeMovies)
	Router.HandleFunc("/movies/upload", movies.UploadMovie)
	Router.HandleFunc("/home", movies.GetHomeMovies)
	Router.HandleFunc("/servers/reload", ReloadServers)
	Router.HandleFunc("/server/reload", ReloadMe)
	Router.HandleFunc("/movies/addserver/{id}", movies.AddServer)
	Router.HandleFunc("/movies/addurl/{id}/{url}", movies.AddUrl)
	Router.HandleFunc("/movies/setserver/{id}/{servername}", movies.SetServer)
	Router.HandleFunc("/tv-show/setserver/{id}/{seasoncode}/{episodecode}/{servername}", movies.SetEpisodeServer)
	Router.HandleFunc("/tv-show/addserver/{id}/{seasoncode}/{episodecode}", movies.AddEpisodeServer)
	Router.HandleFunc("/tv-show/addurl/{id}/{seasoncode}/{episodecode}/{urls}", movies.AddEpisodeUrl)
	Router.HandleFunc("/tv-shows/deleteurls/{id}/{seasoncode}/{episodecode}/{urls}", movies.DeleteUrls)
	Router.HandleFunc("/movies/deleteurls/{id}/{urls}", movies.DeleteUrls)
	Router.HandleFunc("/tv-shows/{id}/addseason", movies.AddSeason)
	Router.HandleFunc("/tv-shows/{id}/{seasonid}/addepisode", movies.AddEpisode)
	Router.HandleFunc("/tv-shows/deleteserver/{id}/{seasoncode}/{episodecode}", movies.DeleteServer)
	Router.HandleFunc("/all/{genre}", movies.GetMoviesByGenre)
	Router.HandleFunc("/{type}/{genre}", movies.GetMoviesByTypeAndGenre)
	Router.HandleFunc("/movies/deleteserver/{id}", movies.DeleteServer)
	Router.HandleFunc("/errors/{package}", Errors)
}

func ReloadServers(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	if req.Method != "GET" {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	servers.ReloadServers()
	res.Write(variables.JsonMarshal(`{"success": true}`))
}


func ReloadMe(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	if req.Method != "GET" {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	movies.Main()
	res.Write(variables.JsonMarshal(`{"success": true}`))
}


func Errors(res http.ResponseWriter, req *http.Request){
	res.Header().Set("content-type", "application/json")
	if req.Method != "GET" {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	params := mux.Vars(req)
	Package := params["package"]
	data, status := variables.GetErrors(Package)
	res.WriteHeader(status)
	res.Write(data)
}