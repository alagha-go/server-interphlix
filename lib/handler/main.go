package handler

import (
	"interphlix/lib/handler/accounts"
	"interphlix/lib/handler/movies"
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
	Router.HandleFunc("/", movies.GetHomeMovies).Methods("GET")
	Router.HandleFunc("/movies/upload", movies.UploadMovie).Methods("POST")
	Router.HandleFunc("/home", movies.GetHomeMovies).Methods("GET")
	Router.HandleFunc("/movies/addserver/{id}", movies.AddServer).Methods("PUT", "UPDATE", "PATCH")
	Router.HandleFunc("/movies/addurl/{id}/{url}", movies.AddUrl).Methods("GET")
	Router.HandleFunc("/movies/setserver/{id}/{servername}", movies.SetServer).Methods("PUT", "UPDATE", "PATCH")
	Router.HandleFunc("/tv-show/setserver/{id}/{seasoncode}/{episodecode}/{servername}", movies.SetEpisodeServer).Methods("PUT", "UPDATE", "PATCH")
	Router.HandleFunc("/tv-show/addserver/{id}/{seasoncode}/{episodecode}", movies.AddEpisodeServer).Methods("PUT", "UPDATE", "PATCH")
	Router.HandleFunc("/tv-show/addurl/{id}/{seasoncode}/{episodecode}/{urls}", movies.AddEpisodeUrl).Methods("PUT", "UPDATE", "PATCH")
	Router.HandleFunc("/tv-shows/deleteurls/{id}/{seasoncode}/{episodecode}/{urls}", movies.DeleteUrls).Methods("DELETE")
	Router.HandleFunc("/movies/deleteurls/{id}/{urls}", movies.DeleteUrls).Methods("DELETE")
	Router.HandleFunc("/tv-shows/{id}/addseason", movies.AddSeason).Methods("POST")
	Router.HandleFunc("/tv-shows/{id}/{seasonid}/addepisode", movies.AddEpisode).Methods("POST")
	Router.HandleFunc("/tv-shows/deleteserver/{id}/{seasoncode}/{episodecode}", movies.DeleteServer).Methods("DELETE")
	Router.HandleFunc("/all/{genre}", movies.GetMoviesByGenre).Methods("GET")
	Router.HandleFunc("/login-url", accounts.LoginUrl).Methods("GET")
	Router.HandleFunc("/login/redirect", accounts.LoginRedirect).Methods("GET")
	Router.HandleFunc("/types", movies.GetTypes).Methods("GET")
	Router.HandleFunc("/myaccount", accounts.GetMyAccount).Methods("GET")
	Router.HandleFunc("/token/refresh", accounts.RenewToken).Methods("GET")
	Router.HandleFunc("/files", accounts.GetMyFiles).Methods("GET")
	Router.HandleFunc("/movies/{type}/{genre}", movies.GetMoviesByTypeAndGenre).Methods("GET")
	Router.HandleFunc("/movies/deleteserver/{id}", movies.DeleteServer).Methods("DELETE")
	Router.HandleFunc("/errors/{package}", Errors).Methods("GET")
}


func Errors(res http.ResponseWriter, req *http.Request){
	res.Header().Set("content-type", "application/json")
	valid := accounts.ValidateRequest(req)
	if !valid {
		res.WriteHeader(http.StatusUnauthorized)
		res.Write(variables.JsonMarshal(variables.Error{Error: "unauthorized"}))
		return
	}
	params := mux.Vars(req)
	Package := params["package"]
	data, status := variables.GetErrors(Package)
	res.WriteHeader(status)
	res.Write(data)
}